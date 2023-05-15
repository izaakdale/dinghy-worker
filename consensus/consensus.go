package consensus

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
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
		Name    string `envconfig:"NAME"`
	}
	StreamLayer struct {
		ln              net.Listener
		serverTLSConfig *tls.Config
		peerTLSConfig   *tls.Config
	}
	snapshot struct {
	}
)

func New(cfg Config) error {
	badgerOpt := badger.DefaultOptions(cfg.DataDir)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		return err
	}

	defer func() {
		if err := badgerDB.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error closing badgerDB: %s\n", err.Error())
		}
	}()

	var raftBinAddr = fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	raftConf := raft.DefaultConfig()
	raftConf.LocalID = raft.ServerID(cfg.Name)
	raftConf.SnapshotThreshold = 1024

	store, err := raftboltdb.NewBoltStore(filepath.Join(cfg.DataDir, "raft.dataRepo"))
	if err != nil {
		return err
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, store)
	if err != nil {
		return err
	}

	snapshotStore, err := raft.NewFileSnapshotStore(cfg.DataDir, raftSnapShotRetain, os.Stdout)
	if err != nil {
		return err
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", raftBinAddr)
	if err != nil {
		return err
	}

	transport, err := raft.NewTCPTransport(raftBinAddr, tcpAddr, maxPool, tcpTimeout, os.Stdout)
	if err != nil {
		return err
	}

	raftServer, err := raft.NewRaft(raftConf, &fsm{}, cacheStore, store, snapshotStore, transport)
	if err != nil {
		return err
	}

	// always start single server as a leader
	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(cfg.Name),
				Address: transport.LocalAddr(),
			},
		},
	}

	raftServer.BootstrapCluster(configuration)
	return nil
}
