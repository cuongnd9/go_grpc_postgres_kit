package v2

import (
	"context"
	"github.com/103cuong/grpc_go_kit/pkg/api/v2"
)

type GreeterServiceServer struct {}

func (s *GreeterServiceServer) SayHello(ctx context.Context, req *v2.HelloRequest) (*v2.HelloResponse, error) {
	return &v2.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
