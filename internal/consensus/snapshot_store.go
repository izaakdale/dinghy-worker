package consensus

import (
	"io"

	"github.com/hashicorp/raft"
)

type snapshotStore struct {
}

func (s *snapshotStore) Create(version raft.SnapshotVersion, index, term uint64, configuration raft.Configuration, configurationIndex uint64, trans raft.Transport) (raft.SnapshotSink, error) {
	return &snapshotSink{}, nil
}

// List is used to list the available snapshots in the store.
// It should return then in descending order, with the highest index first.
func (s *snapshotStore) List() ([]*raft.SnapshotMeta, error) {
	return []*raft.SnapshotMeta{}, nil
}

// Open takes a snapshot ID and provides a ReadCloser. Once close is
// called it is assumed the snapshot is no longer needed.
func (s *snapshotStore) Open(id string) (*raft.SnapshotMeta, io.ReadCloser, error) {
	return &raft.SnapshotMeta{}, nil, nil
}
