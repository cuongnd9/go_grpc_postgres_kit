package server

import (
	"context"
	"github.com/103cuong/grpc_go_kit/pkg/api/server"
)

type GreeterServiceServer struct {}

func (s *GreeterServiceServer) SayHello(ctx context.Context, req *server.HelloRequest) (*server.HelloResponse, error) {
	return &server.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
