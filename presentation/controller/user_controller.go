package controller

import (
	"go-clean-app/application/usecase"
	"go-clean-app/presentation/request"
	"go-clean-app/presentation/response"
	"net/http"
	"strconv"

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

func (c *UserController) FindByID(ctx *gin.Context) {
	req := request.UserFindByID{
		ID: ctx.Param("id"),
	}
	err := req.Validate()
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}
	user, err := c.usecase.FindByID(ctx, id)
	response.User(ctx, user, err)
}
