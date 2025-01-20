package helpers

import (
	"go-clean-app/config"
	"go-clean-app/infrastructure/driver"
	"log"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func Initialize(envPath string) (*gin.Context, *driver.ShardingManager) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	err := config.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
		return nil, nil
	}
	db := driver.Initialize()
	return ctx, db
}
