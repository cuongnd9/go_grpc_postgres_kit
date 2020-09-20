package configs

import (
	"github.com/mcuadros/go-defaults"
	"os"
)

type AppConfig struct {
	GoENV string `default:"development"`
	Port string `default:"50000"`
}

var AppConf AppConfig

func BuildAppConfig() {
	AppConf = AppConfig{
		GoENV: os.Getenv("GO_ENV"),
		Port:  os.Getenv("PORT"),
	}
	defaults.SetDefaults(&AppConf)
}
