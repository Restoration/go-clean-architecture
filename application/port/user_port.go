package port

import (
	"go-clean-app/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserPort interface {
	FindAll(ctx *gin.Context, db *gorm.DB) (domain.Users, error)
}
