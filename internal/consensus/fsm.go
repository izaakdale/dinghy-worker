package consensus

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"

	"github.com/hashicorp/raft"
	v1 "github.com/izaakdale/dinghy-worker/api/v1"
	"github.com/izaakdale/dinghy-worker/internal/store"
	"google.golang.org/protobuf/proto"
)

type fsm struct {
	client *store.Client
	logs   []*raft.Log
}

func (f *fsm) Apply(record *raft.Log) any {
	f.logs = append(f.logs, record)
	switch record.Type {
	case raft.LogCommand:
		var req v1.InsertRequest
		if err := proto.Unmarshal(record.Data, &req); err != nil {
			return err
		}

		// if no value was present this is a delete request
		if req.Value == "" {
			var req v1.DeleteRequest
			if err := proto.Unmarshal(record.Data, &req); err != nil {
				return err
			}

			if err := f.client.Delete([]byte(req.Key)); err != nil {
				return err
			}
			return nil
		}

		if err := f.client.Insert([]byte(req.Key), []byte(req.Value)); err != nil {
			return err
		}
	}
	return nil
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	return &snapshot{f}, nil
}

func (f *fsm) Restore(snapshot io.ReadCloser) error {
	log.Printf("attempted restore in fsm")
	var logs []*raft.Log
	if err := gob.NewDecoder(snapshot).Decode(&logs); err != nil {
		return err
	}

	for _, l := range logs {
		log.Printf("applying log: %d\n", l.Index)
		f.Apply(l)
	}

	return nil
}

type snapshot struct {
	fsm *fsm
}

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	log.Printf("attempting persist in fsm")
	var encodedLogs bytes.Buffer
	if err := gob.NewEncoder(&encodedLogs).Encode(s.fsm.logs); err != nil {
		return err
	}
	sink.Write(encodedLogs.Bytes())
	log.Printf("persisted\n")
	sink.Close()
	return nil
}

func (s *snapshot) Release() {
	s.fsm.logs = make([]*raft.Log, 0)
	log.Printf("attempted release in fsm")
}
