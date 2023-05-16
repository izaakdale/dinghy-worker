package consensus

import (
	"io"
	"log"

	"github.com/hashicorp/raft"
	v1 "github.com/izaakdale/dinghy-worker/api/v1"
	"github.com/izaakdale/dinghy-worker/internal/store"
	"google.golang.org/protobuf/proto"
)

type RequestType uint8

const AppendRequestType RequestType = 0

type fsm struct {
	client *store.Client
}

type ApplyResponse struct {
}

func (f *fsm) Apply(record *raft.Log) any {
	switch record.Type {
	case raft.LogCommand:
		var req v1.InsertRequest
		if err := proto.Unmarshal(record.Data, &req); err != nil {
			return err
		}
		if err := f.client.Insert([]byte(req.Key), []byte(req.Value)); err != nil {
			return err
		}
	}
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
