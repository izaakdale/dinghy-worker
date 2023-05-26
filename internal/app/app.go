package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/serf/serf"
	v1 "github.com/izaakdale/dinghy-worker/api/v1"
	"github.com/izaakdale/dinghy-worker/internal/consensus"
	"github.com/izaakdale/dinghy-worker/internal/discovery"
	"github.com/izaakdale/dinghy-worker/internal/server"
	"github.com/izaakdale/dinghy-worker/internal/store"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

var spec Specification

type Specification struct {
	GRPCAddr     string `envconfig:"GRPC_ADDR"`
	GRPCPort     int    `envconfig:"GRPC_PORT"`
	Name         string `envconfig:"NAME"`
	discoveryCfg discovery.Config
	consensusCfg consensus.Config
}

type App struct{}

func New() *App {
	if err := envconfig.Process("", &spec); err != nil {
		log.Fatalf("failed to process env vars: %v", err)
	}
	if err := envconfig.Process("", &spec.discoveryCfg); err != nil {
		log.Fatalf("failed to process discovery env vars: %v", err)
	}
	if err := envconfig.Process("", &spec.consensusCfg); err != nil {
		log.Fatalf("failed to process consensus env vars: %v", err)
	}
	return &App{}
}

func (a *App) Run() {
	log.Printf("hello, my name is %s\n", spec.Name)

	badgerOpt := badger.DefaultOptions(spec.consensusCfg.DataDir)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		log.Fatalf("failed to start up badger db: %v", err)
	}

	defer func() {
		if err := badgerDB.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error closing badgerDB: %s\n", err.Error())
		}
	}()

	dbClient, err := store.New(badgerDB)
	if err != nil {
		log.Fatalf("failed to start up store client: %v", err)
	}

	raftNode, err := consensus.New(spec.Name, spec.consensusCfg, dbClient)
	if err != nil {
		log.Fatalf("failed to start up raft: %v", err)
	}

	gAddr := fmt.Sprintf("%s:%d", spec.GRPCAddr, spec.GRPCPort)
	ln, err := net.Listen("tcp", gAddr)
	if err != nil {
		log.Fatalf("failed to start up grpc listener: %v", err)
	}

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	srv := server.New(spec.Name, raftNode, dbClient)
	v1.RegisterWorkerServer(gsrv, srv)

	errCh := make(chan error)
	go func(ch chan error) {
		ch <- gsrv.Serve(ln)
	}(errCh)

	raftAddr := fmt.Sprintf("%s:%d", spec.consensusCfg.Addr, spec.consensusCfg.Port)

	serfNode, evCh, err := discovery.NewMembership(spec.Name, spec.discoveryCfg, discovery.Tag{
		Key:   "grpc_addr",
		Value: gAddr,
	}, discovery.Tag{
		Key:   "raft_addr",
		Value: raftAddr,
	})
	defer serfNode.Leave()
	if err != nil {
		log.Fatal(err)
	}

	// let start up to settle before sending state heartbeats
	t := newDelayedTicker()

	shCh := make(chan os.Signal, 2)
	signal.Notify(shCh, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-shCh:
			err := serfNode.Leave()
			if err != nil {
				log.Fatalf("error leaving cluster %v", err)
			}
			if raftNode.State() == raft.Leader {
				f := raftNode.RemoveServer(raft.ServerID(spec.Name), 0, 0)
				if f.Error() != nil {
					log.Printf("error leaving raft cluster: %v\n", f.Error())
				}
			}
			os.Exit(1)
		case e := <-evCh:
			if raftNode.State() == raft.Leader {
				switch e.EventType() {
				case serf.EventMemberLeave, serf.EventMemberFailed:
					for _, m := range e.(serf.MemberEvent).Members {
						name, ok := m.Tags["name"]
						if !ok {
							log.Printf("no name tag in leaving member\n")
						}
						raftNode.RemoveServer(raft.ServerID(name), 0, 0)
					}
				}
			}
		case <-t:
			log.Printf("ticker triggered\n")
			log.Printf("state: %s\n", raftNode.State())
			if raftNode.State() == raft.Leader {
				payload := v1.LeaderHeaderbeat{
					Name:      spec.Name,
					RaftState: raftNode.State().String(),
					GrpcAddr:  gAddr,
					RaftAddr:  raftAddr,
				}
				bytes, err := proto.Marshal(&payload)
				if err != nil {
					log.Printf("error marshalling leader heartbeat payload: %v\n", err)
				}
				serfNode.UserEvent("leader-notification", bytes, true)
			}
		case err := <-errCh:
			log.Fatalf("grpc server errored: %v", err)
		}
	}
}

func newDelayedTicker() <-chan time.Time {
	time.Sleep(10 * time.Second)
	return time.Tick(time.Second)
}
