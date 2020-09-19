package main

import (
	"fmt"
	"github.com/103cuong/grpc_go_kit/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "global error: %v\n", err)
		os.Exit(1)
	}
}
