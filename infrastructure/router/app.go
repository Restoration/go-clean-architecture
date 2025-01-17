package router

import (
	"go-clean-app/config"
	"go-clean-app/di"
	"go-clean-app/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func App(router *gin.Engine, db *gorm.DB) {
	apiConfig := config.GetAPIConfig()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.OpenTelemetryMiddleware())
	version := "/" + apiConfig.ApiVersion
	v1 := router.Group(version)
	v1.GET("/users", di.DiUser(db).Users)
}
