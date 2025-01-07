package router

import (
	"go-clean-app/di"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func App(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", di.DiUser(db).Users)
}
