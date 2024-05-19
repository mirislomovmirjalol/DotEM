package store

type IStore interface {
	Set(bucketName string, key string, value []byte) error
	Get(bucketName string, key string) ([]byte, error)
	Delete(bucketName string, key string) error
	GetKeys(bucketName string) ([]string, error)
}
