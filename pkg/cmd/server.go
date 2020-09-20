package cmd

import (
	"context"
	"github.com/103cuong/go_grpc_postgres_kit/configs"
	"github.com/103cuong/go_grpc_postgres_kit/internal/migration"
	"github.com/103cuong/go_grpc_postgres_kit/pkg/protocol"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func RunServer() error {
	ctx := context.Background()

	var err error
	configs.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  configs.DatabaseDSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln("connect to database failed")
		panic("connect to database failed")
	}

	// migrate database.
	migration.MigrateDB("up", "./internal/migration")

	//  grpc client registration
	defer RegisterClient()

	return protocol.RunServer(ctx, configs.AppConf.Port)
}
