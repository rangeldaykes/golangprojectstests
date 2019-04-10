package repositorycache

import "store/infrastucture/data/infraredis"

// RepositoryFcvs methods return fcvs data from redis
type RepositoryCacheUser struct {
	infraredis.IInfraPersistenceRedis
}

func (rb RepositoryCacheUser) GetKey(key string) (string, error) {
	return rb.IInfraPersistenceRedis.GetKey(key)
}

// GetkeyGeneric get any key from redis
// func (rf RepositoryFcvs) GetkeyGeneric(key string) (string, error) {
// 	return rf.GetKey(key)
// }
