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
	// ファイルの署名処理
	// 完了後Spanに記録
	span.SetAttributes(
		attribute.String("aws.service", "s3"),
		attribute.String("aws.operation", "Presigned URL"),
		attribute.String("aws.region", "us-east-1"),
		attribute.String("aws.request_id", "XXXXXXXXXXXX"),
		attribute.String("aws.bucket", bucket),
		attribute.Int("aws.storageKey", storageKey),
	)
	return nil, nil
}
