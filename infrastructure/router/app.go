package router

import (
	"go-clean-app/di"
	"go-clean-app/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func App(router *gin.Engine, db *gorm.DB) {
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.OpenTelemetryMiddleware())
	router.GET("/users", di.DiUser(db).Users)
}
