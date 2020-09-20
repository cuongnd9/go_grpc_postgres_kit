package cmd

import (
	"context"
	"github.com/103cuong/grpc_go_kit/configs"
	"github.com/103cuong/grpc_go_kit/internal/migration"
	"github.com/103cuong/grpc_go_kit/pkg/protocol"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunServer() error {
	ctx := context.Background()

	var err error
	configs.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  configs.BuildDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("connect to database failed")
	}

	// migrate database.
	migration.MigrateDB("up", "./internal/migration")

	return protocol.RunServer(ctx, "50000")
}
