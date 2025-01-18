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
	ApiVersion string `envconfig:"API_VERSION" required:"true"`
	Port       int    `envconfig:"PORT" required:"true"`
	HostURL    string `envconfig:"HOST_URL" required:"true"`
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

type DB1 struct {
	Host     string `envconfig:"POSTGRES_HOST_1" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT_1" required:"true"`
	User     string `envconfig:"POSTGRES_USER_1" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD_1" required:"true"`
	Name     string `envconfig:"POSTGRES_DATABASE_1" required:"true"`
}

type DB2 struct {
	Host     string `envconfig:"POSTGRES_HOST_2" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT_2" required:"true"`
	User     string `envconfig:"POSTGRES_USER_2" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD_2" required:"true"`
	Name     string `envconfig:"POSTGRES_DATABASE_2" required:"true"`
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

func GetDB1Config() *DB {
	return config.DB
}

func GetDB2Config() *DB {
	return config.DB
}
