package di

import (
	"go-clean-app/application/port"
	"go-clean-app/infrastructure/gateway"
)

func DiAWSGateway() port.AWSPort {
	return gateway.NewStubAWSGateway()
}
