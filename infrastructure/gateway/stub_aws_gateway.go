package gateway

type StubAWSGateway struct{}

func NewStubAWSGateway() *StubAWSGateway {
	return &StubAWSGateway{}
}

func (repo *StubAWSGateway) CreatePreSignedURL(bucket string, storageKey int) (url *string, err error) {
	return nil, nil
}
