package protocol

import (
	"context"
	"fmt"
	API "github.com/103cuong/go_grpc_postgres_kit/pkg/api/server"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, port string) error { // optional param: db
	// listen on the local network
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	// register service
	grpcServer := grpc.NewServer()
	API.RegisterGreeterServer(grpcServer, &server.GreeterServer{})
	API.RegisterCatServiceServer(grpcServer, &server.CatServer{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("💤 shutting down gRPC server")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("💅 server ready at 0.0.0.0:%s", port)
	return grpcServer.Serve(listen)
}
