package interfaces

// IRepositoryRedisBase reids basic interface
type IRepositoryRedisBase interface {
	GetKey(key string) (string, error)
}
