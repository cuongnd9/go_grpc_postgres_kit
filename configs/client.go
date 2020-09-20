package configs

import (
	"github.com/103cuong/go_grpc_postgres_kit/pkg/api/client"
	"github.com/mcuadros/go-defaults"
	"os"
)

type ClientConfig struct {
	GreeterClientPort string `default:"50001"`
}

var ClientConf ClientConfig
var GreeterClient client.GreeterClient

func BuildClientConfig()  {
	ClientConf = ClientConfig{GreeterClientPort: os.Getenv("GREETER_CLIENT_PORT")}
	defaults.SetDefaults(&ClientConf)
}
