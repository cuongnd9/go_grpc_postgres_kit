package main

import (
	"fmt"
	"os"

	"github.com/103cuong/go_grpc_postgres_kit/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "global error: %v\n", err)
		os.Exit(1)
	}
}
