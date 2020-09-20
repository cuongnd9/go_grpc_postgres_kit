package server

import (
	"context"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/api/server"
)

type GreeterServer struct {}

func (s *GreeterServer) SayHello(ctx context.Context, req *server.HelloRequest) (*server.HelloResponse, error) {
	return &server.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
