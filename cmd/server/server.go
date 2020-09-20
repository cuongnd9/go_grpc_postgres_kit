package main

import (
	"fmt"
	"github.com/103cuong/go_grpc_postgres_kit/configs"
	"os"

	"github.com/103cuong/go_grpc_postgres_kit/pkg/cmd"
)

func main() {
	configs.BuildConfig()

	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "global error: %v\n", err)
		os.Exit(1)
	}
}
