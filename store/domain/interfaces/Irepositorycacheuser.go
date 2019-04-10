package interfaces

// IRepositoryFcvs provides redis iterations to fcvs
type IRepositoryCacheUser interface {
	//GetkeyGeneric(key string) (string, error)
	GetKey(key string) (string, error)
	//infraredis.IInfraPersistenceRedis // IRepositoryRedisBase
}
