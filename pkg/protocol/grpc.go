package protocol

import (
	"context"
	"fmt"
	greeterApi "github.com/103cuong/grpc_go_kit/pkg/api/server"
	greeterServer "github.com/103cuong/grpc_go_kit/pkg/grpc/server"
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
	server := grpc.NewServer()
	greeterApi.RegisterGreeterServer(server, &greeterServer.GreeterServiceServer{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("💤 shutting down gRPC server")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("💅 server ready at 0.0.0.0:%s", port)
	return server.Serve(listen)
}
