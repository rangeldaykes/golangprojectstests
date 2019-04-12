package repositorycache

import "store/infrastucture/data/infraredis"

type RepositoryCacheUser struct {
	infraredis.IInfraPersistenceRedis
}

func (rb RepositoryCacheUser) GetUser(ID string) (string, error) {
	return rb.IInfraPersistenceRedis.GetKey(ID)
}
