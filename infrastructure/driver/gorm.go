package driver

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go-clean-app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func postgresDSN(config *config.DB) string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, strconv.Itoa(config.Port), config.Name)
	return dsn
}

func setupPostgreSQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		return nil, fmt.Errorf("connection error")
	}
	sqlDB, _ := db.DB()
	config := config.GetDBConfig()
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	return db, nil
}

func Initialize() *gorm.DB {
	cfg := config.GetDBConfig()
	// cfg1 := config.GetDB1Config()
	// cfg2 := config.GetDB2Config()
	conn, err := setupPostgreSQL(postgresDSN(cfg))
	if err != nil {
		panic(err)
	}
	if err := conn.Error; err != nil {
		panic(err)
	}
	conn.Logger = logger.Default.LogMode(logger.Silent)
	return conn
}
