package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/izaakdale/dinghy-worker/consensus"
	"github.com/izaakdale/dinghy-worker/discovery"
	"github.com/kelseyhightower/envconfig"
)

var spec Specification

type Specification struct {
	GRPCAddr     string `envconfig:"GRPC_ADDR"`
	GRPCPort     int    `envconfig:"GRPC_PORT"`
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
	// raftNode, err := consensus.New(spec.consensusCfg)
	_, err := consensus.New(spec.consensusCfg)
	if err != nil {
		log.Fatalf("failed to start up raft: %v", err)
	}

	// set up a grpc server here to handle db transactions, pass raftNode to use raftNode.Apply(grpcMsg)

	serfNode, evCh, err := discovery.NewMembership(spec.discoveryCfg, discovery.Tag{
		Key:   "grpc_addr",
		Value: fmt.Sprintf("%s:%d", spec.GRPCAddr, spec.GRPCPort),
	})
	defer serfNode.Leave()
	if err != nil {
		log.Fatal(err)
	}

	shCh := make(chan os.Signal, 2)
	signal.Notify(shCh, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-shCh:
			err := serfNode.Leave()
			if err != nil {
				log.Fatalf("error leaving cluster %v", err)
			}
			os.Exit(1)
		case <-evCh:
			log.Println("event channel triggered")
		}
	}
}
