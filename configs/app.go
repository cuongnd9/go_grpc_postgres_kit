package configs

import (
	"github.com/mcuadros/go-defaults"
	"os"
)

type AppConfig struct {
	GoENV string `default:"development"`
	Port string `default:"50000"`
}

var App AppConfig

func BuildAppConfig() {
	App = AppConfig{
		GoENV: os.Getenv("GO_ENV"),
		Port:  os.Getenv("PORT"),
	}
	defaults.SetDefaults(&App)
}
