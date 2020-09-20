package configs

import (
	"os"

	"fmt"
	"github.com/mcuadros/go-defaults"
	"gorm.io/gorm"
)

// DB GORM DB instance
var DB *gorm.DB

// DatabaseConfig database configuration
type DatabaseConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     string `default:"5432"`
	User     string `default:"postgres"`
	DBName   string `default:"postgres"`
	Password string `default:"postgres"`
}

var Database DatabaseConfig
var DatabaseDSN string

// BuildDatabaseConfig build DB config
func BuildDatabaseConfig() {
	Database = DatabaseConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DBNAME"),
		Password: os.Getenv("PG_PASSWORD"),
	}
	defaults.SetDefaults(&Database)
	DatabaseDSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Database.Host,
		Database.User,
		Database.Password,
		Database.DBName,
		Database.Port,
	)
}

