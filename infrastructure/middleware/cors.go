package middleware

import (
	"time"

	"go-clean-app/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := config.GetAPIConfig()
	return cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			// "X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		AllowOrigins: []string{
			config.Host,
		},
		AllowOriginFunc: func(origin string) bool {
			return origin == config.Host
		},
		MaxAge: 24 * time.Hour,
	})
}
