package main

import "github.com/103cuong/grpc_go_kit/internal/migration"

func main()  {
	migration.MigrateDB("down", "./internal/migration")
}
