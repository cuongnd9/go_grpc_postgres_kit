package cmd

import (
	"context"
	"github.com/103cuong/grpc_go_kit/pkg/protocol/grpc"
)

func RunServer() error {
	ctx := context.Background()

	return grpc.RunServer(ctx, "50000")
}
