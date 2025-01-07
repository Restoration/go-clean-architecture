package response

import (
	"go-clean-app/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Users(ctx *gin.Context, users domain.Users, err error) {
	if err != nil {
		Error(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
