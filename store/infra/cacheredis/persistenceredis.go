package cacheredis

// IInfraPersistenceRedis methos to works with redis
type IInfraPersistenceRedis interface {
	Ping() (string, error)
	GetKey(key string) (string, error)
}
