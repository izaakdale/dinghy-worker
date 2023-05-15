package consensus

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type RequestType uint8

const AppendRequestType RequestType = 0

type fsm struct{}

func (f *fsm) Apply(record *raft.Log) any {
	return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	log.Printf("attempted snapshot in fsm")
	return nil, nil
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	log.Printf("attempted persist in fsm")
	return nil
}

func (s *snapshot) Release() {
	log.Printf("attempted release in fsm")
}

func (f *fsm) Restore(snapshot io.ReadCloser) error {
	log.Printf("attempted restore in fsm")
	return nil
}
