package configs

import (
	"gorm.io/gorm"
	"os"

	"fmt"
	"github.com/mcuadros/go-defaults"
)

// DatabaseConfig database configuration
type DatabaseConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     string `default:"5432"`
	User     string `default:"postgres"`
	DBName   string `default:"postgres"`
	Password string `default:"postgres"`
}

// DB GORM DB instance
var DB *gorm.DB
var DatabaseConf DatabaseConfig
var DatabaseDSN string

// BuildDatabaseConfig build DB config
func BuildDatabaseConfig() {
	DatabaseConf = DatabaseConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DBNAME"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	defaults.SetDefaults(&DatabaseConf)
	DatabaseDSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DatabaseConf.Host,
		DatabaseConf.User,
		DatabaseConf.Password,
		DatabaseConf.DBName,
		DatabaseConf.Port,
	)
}

