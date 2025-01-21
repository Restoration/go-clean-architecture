package router

import (
	"go-clean-app/config"
	"go-clean-app/di"
	"go-clean-app/infrastructure/driver"
	"go-clean-app/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func App(router *gin.Engine, db *driver.ShardingManager) {
	apiConfig := config.GetAPIConfig()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.OpenTelemetryMiddleware())
	version := "/" + apiConfig.ApiVersion
	v1 := router.Group(version)
	v1.GET("/users", di.DiUser(db).Users)
	v1.GET("/user:id", di.DiUser(db).FindByID)
}
