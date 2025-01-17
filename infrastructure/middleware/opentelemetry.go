package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func OpenTelemetryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tracer := otel.Tracer("go-clean-app")
		// トレースの開始
		ctx, span := tracer.Start(ctx, "http.request",
			trace.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.url", c.Request.URL.String()),
				attribute.String("http.client_ip", c.ClientIP()),
				attribute.Int("http.status_code", c.Writer.Status()),
			),
		)
		defer span.End()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
