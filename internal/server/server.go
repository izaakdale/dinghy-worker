package server

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/raft"
	v1 "github.com/izaakdale/dinghy-worker/api/v1"
	"github.com/izaakdale/dinghy-worker/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// ensure our server adheres to grpc worker server
var _ v1.WorkerServer = (*Server)(nil)

type Server struct {
	v1.UnimplementedWorkerServer
	name      string
	consensus *raft.Raft
	client    store.Transactioner
}

func New(name string, consensus *raft.Raft, c *store.Client) *Server {
	return &Server{
		name:      name,
		consensus: consensus,
		client:    c,
	}
}

func (s *Server) Join(ctx context.Context, request *v1.JoinRequest) (*v1.JoinResponse, error) {
	if s.consensus.State() != raft.Leader {
		return nil, fmt.Errorf("must make this request to leader")
	}

	future := s.consensus.AddVoter(raft.ServerID(request.ServerId), raft.ServerAddress(request.ServerAddr), 0, 0)
	if future.Error() != nil {
		return nil, fmt.Errorf("error adding voter to raft cluster: %v", future.Error())
	}
	return &v1.JoinResponse{}, nil
}

func (s *Server) Insert(ctx context.Context, request *v1.InsertRequest) (*v1.InsertResponse, error) {
	if s.consensus.State() != raft.Leader {
		return nil, fmt.Errorf("cannot insert since I'm not the raft leader.")
	}
	protoBytes, err := proto.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request to proto bytes: %v", err)
	}

	applyFuture := s.consensus.Apply(protoBytes, 500*time.Millisecond)
	if err := applyFuture.Error(); err != nil {
		return nil, fmt.Errorf("error applying request to raft cluster: %v", err)
	}

	if err, ok := applyFuture.Response().(error); ok {
		return nil, fmt.Errorf("error from raft node when applying request: %v", err)
	}
	return &v1.InsertResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, request *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Fetch(ctx context.Context, request *v1.FetchRequest) (*v1.FetchResponse, error) {
	val, err := s.client.Fetch([]byte(request.Key))
	if err != nil {
		if err == store.ErrNotFound {
			st := status.New(codes.NotFound, "no records exist for the given key.")
			return nil, st.Err()
		}
		return nil, err
	}
	return &v1.FetchResponse{Key: request.Key, Value: string(val)}, nil
}

func (s *Server) RaftState(ctx context.Context, request *v1.RaftStateRequest) (*v1.RaftStateResponse, error) {
	return &v1.RaftStateResponse{
		State: s.consensus.Stats()["state"],
	}, nil
}
