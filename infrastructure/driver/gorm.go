package driver

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go-clean-app/config"
	"go-clean-app/infrastructure/tracer"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//	func postgresDSN(config *config.DB) string {
//		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, strconv.Itoa(config.Port), config.Name)
//		return dsn
//	}

func postgresDSN(user, password, host, port, dbName string) string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)
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

type ShardingManager struct {
	dbShards map[int]*gorm.DB
}

func NewShardingManager(shardConfigs map[int]string) (*ShardingManager, error) {
	dbShards := make(map[int]*gorm.DB)
	for shardID, dsn := range shardConfigs {
		conn, err := setupPostgreSQL(dsn)
		if err != nil {
			panic(err)
		}
		if err := conn.Error; err != nil {
			panic(err)
		}
		conn.Logger = logger.Default.LogMode(logger.Silent)
		dbShards[shardID] = conn
		tracer.RegisterGORMCallbacks(conn)

	}
	return &ShardingManager{dbShards: dbShards}, nil
}

func (sm *ShardingManager) GetShardID(userID int) int {
	return (userID % len(sm.dbShards)) + 1
}

func (sm *ShardingManager) GetDBForUser(userID int) *gorm.DB {
	shardID := sm.GetShardID(userID)
	return sm.dbShards[shardID]
}

func (sm *ShardingManager) GetShards() map[int]*gorm.DB {
	return sm.dbShards
}

func Initialize() *ShardingManager {
	cfg := config.GetDBConfig()
	cfg1 := config.GetDB1Config()
	cfg2 := config.GetDB2Config()
	shardConfigs := map[int]string{
		1: postgresDSN(cfg.User, cfg.Password, cfg.Host, strconv.Itoa(cfg.Port), cfg.Name),
		2: postgresDSN(cfg1.User, cfg1.Password, cfg1.Host, strconv.Itoa(cfg1.Port), cfg1.Name),
		3: postgresDSN(cfg2.User, cfg2.Password, cfg2.Host, strconv.Itoa(cfg2.Port), cfg2.Name),
	}
	shardingManager, err := NewShardingManager(shardConfigs)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize sharding manager: %v", err))
	}
	return shardingManager
}
