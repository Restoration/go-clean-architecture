package tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func RegisterGORMCallbacks(db *gorm.DB) {
	// Query開始前にトレーススパンを開始
	db.Callback().Query().Before("gorm:query").Register("otel:before_query", beforeQuery)
	// Query終了後にトレーススパンを終了
	db.Callback().Query().After("gorm:query").Register("otel:after_query", afterQuery)
}

var tracer = otel.Tracer("gorm-tracer")

func beforeQuery(db *gorm.DB) {
	ctx, span := tracer.Start(db.Statement.Context, "gorm.query",
		trace.WithAttributes(
			attribute.String("db.system", "postgresql"),
			attribute.String("db.statement", db.Statement.SQL.String()), // 実行するSQL
		),
	)
	db.Statement.Context = ctx
	db.InstanceSet("otel:span", span)
}

func afterQuery(db *gorm.DB) {
	spanInterface, ok := db.InstanceGet("otel:span")
	if !ok {
		return
	}

	span := spanInterface.(trace.Span)
	span.SetAttributes(
		attribute.Int("db.rows_affected", int(db.RowsAffected)), // 影響を受けた行数
		attribute.String("db.error", db.Error.Error()),          // エラー情報（ある場合）
	)
	span.End()
}
