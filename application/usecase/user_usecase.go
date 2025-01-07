package usecase

import (
	"go-clean-app/domain"

	"github.com/gin-gonic/gin"
)

type UserUseCase interface {
	FindAll(ctx *gin.Context) (domain.Users, error)
}
