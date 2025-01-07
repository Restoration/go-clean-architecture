package port

type AWSPort interface {
	CreatePreSignedURL(bucket string, storageKey int) (url *string, err error)
}
