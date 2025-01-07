package controller

import (
	"go-clean-app/application/usecase"
	"go-clean-app/presentation/response"

	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	Users(ctx *gin.Context)
}
type UserController struct {
	usecase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) *UserController {
	return &UserController{usecase: usecase}
}

func (c *UserController) Users(ctx *gin.Context) {
	users, err := c.usecase.FindAll(ctx)
	response.Users(ctx, users, err)
}
