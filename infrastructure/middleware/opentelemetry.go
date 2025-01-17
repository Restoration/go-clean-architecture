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
		// エンドポイントパス
		path := c.FullPath()
		if path == "" {
			path = "unknown" // パスが取得できない場合のデフォルト値
		}
		// トレースの開始
		ctx, span := tracer.Start(ctx, "http.request",
			trace.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.url", c.Request.URL.String()),
				attribute.String("http.client_ip", c.ClientIP()),
				attribute.String("http.path", path),
			),
		)
		defer span.End()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		// ステータスコードの記録は c.Next() 実行後
		span.SetAttributes(
			attribute.Int("http.status_code", c.Writer.Status()),
		)
	}
}
