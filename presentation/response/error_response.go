package response

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Error(ctx *gin.Context, statusCode int, err error) {
	ctx.Error(err)
	ctx.JSON(statusCode, gin.H{
		"message": errors.Cause(err).Error(),
	})
	ctx.Abort()
}
