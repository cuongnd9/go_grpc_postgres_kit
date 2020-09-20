package cmd

import (
	"fmt"
	"github.com/103cuong/go_grpc_postgres_kit/configs"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/api/client"
	"google.golang.org/grpc"
	"log"
)

func buildGRPCConnection(port string) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection gRPC failed: %v\n", err)
	}
	defer conn.Close()
	return conn
}

func RegisterClient() {
	greeterClientConn := buildGRPCConnection(configs.ClientConf.GreeterClientPort)
	configs.GreeterClient = client.NewGreeterClient(greeterClientConn)
}
