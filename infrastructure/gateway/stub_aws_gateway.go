package gateway

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type StubAWSGateway struct{}

func NewStubAWSGateway() *StubAWSGateway {
	return &StubAWSGateway{}
}

func (repo *StubAWSGateway) CreatePreSignedURL(bucket string, storageKey int) (url *string, err error) {
	tracer := otel.Tracer("go-clean-app.gateway")
	_, span := tracer.Start(context.Background(), "StubAWSGateway.CreatePreSignedURL")
	defer span.End()
	// 何らかの処理
	// 完了後Spanに記録
	span.SetAttributes(
		attribute.String("aws.service", "s3"),
		attribute.String("aws.operation", "Presigned URL"),
		attribute.String("aws.bucket", bucket),
		attribute.Int("aws.storageKey", storageKey),
	)
	return nil, nil
}
