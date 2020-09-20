package main

import "github.com/103cuong/go_grpc_postgres_kit/internal/migration"

func main() {
	migration.MigrateDB("down", "./internal/migration")
}
