package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API *API
	DB  *DB
}

type API struct {
	Port int    `envconfig:"PORT" required:"true"`
	Host string `envconfig:"HOST_URL" required:"true"`
}

type DB struct {
	Host            string `envconfig:"POSTGRES_HOST" required:"true"`
	Port            int    `envconfig:"POSTGRES_PORT" required:"true"`
	User            string `envconfig:"POSTGRES_USER" required:"true"`
	Password        string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Name            string `envconfig:"POSTGRES_DATABASE" required:"true"`
	IsSSL           string `envconfig:"DB_USE_SSL" default:"disable"`
	MaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS" default:"255"`
	MaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"255"`
	ConnMaxLifetime int    `envconfig:"DB_CONN_MAX_LIFETIME" default:"255"`
}

var config Config
var envPath = fmt.Sprintf("./%s.env", os.Getenv("GO_ENV"))

func Load(filePath string) error {
	if filePath != "" {
		envPath = filePath
	}
	if _, err := os.Stat(envPath); err == nil {
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("failed to load %s, %v\n", envPath, err)
		}
	}
	if err := envconfig.Process("", &config); err != nil {
		return err
	}
	return nil
}

func GetAPIConfig() *API {
	return config.API
}

func GetDBConfig() *DB {
	return config.DB
}
