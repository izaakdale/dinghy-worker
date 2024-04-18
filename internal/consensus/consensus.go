package consensus

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/izaakdale/dinghy-worker/internal/store"
)

const (
	// The maxPool controls how many connections we will pool.
	maxPool = 3

	// The timeout is used to apply I/O deadlines. For InstallSnapshot, we multiply
	// the timeout by (SnapshotSize / TimeoutScale).
	// https://github.com/hashicorp/raft/blob/v1.1.2/net_transport.go#L177-L181
	tcpTimeout = 10 * time.Second

	// The `retain` parameter controls how many
	// snapshots are retained. Must be at least 1.
	raftSnapShotRetain = 2

	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

type (
	Config struct {
		Addr    string `envconfig:"RAFT_ADDR"`
		Port    int    `envconfig:"RAFT_PORT"`
		DataDir string `envconfig:"DATA_DIR"`
	}
)

func New(name string, cfg Config, dbClient *store.Client) (*raft.Raft, error) {
	var raftBinAddr = fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	raftConf := raft.DefaultConfig()
	raftConf.LocalID = raft.ServerID(name)
	raftConf.SnapshotThreshold = 1024

	bolt, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raft.dataRepo"))
	if err != nil {
		return nil, err
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, bolt)
	if err != nil {
		return nil, err
	}

	snapshotStore, err := raft.NewFileSnapshotStore(cfg.DataDir, raftSnapShotRetain, os.Stdout)
	if err != nil {
		return nil, err
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", raftBinAddr)
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(raftBinAddr, tcpAddr, maxPool, tcpTimeout, os.Stdout)
	if err != nil {
		return nil, err
	}

	// TODO shouldn't be using bolt as stable store here. This should be something like s3
	// raftServer, err := raft.NewRaft(raftConf, &fsm{dbClient}, cacheStore, bolt, snapshotStore, transport)
	raftServer, err := raft.NewRaft(raftConf, &fsm{dbClient}, cacheStore, bolt, snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	// always start single server as a leader
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(name),
				Address: transport.LocalAddr(),
			},
		},
	}

	raftServer.BootstrapCluster(configuration)
	return raftServer, nil
}
