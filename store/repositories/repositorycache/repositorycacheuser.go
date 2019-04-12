package repositorycache

import "store/infrastucture/data/infraredis"

type repositoryCacheUser struct {
	infraredis.IInfraPersistenceRedis
}

func NewRepositoryCacheUser(infra infraredis.IInfraPersistenceRedis) *repositoryCacheUser {
	return &repositoryCacheUser{infra}
}

func (rb repositoryCacheUser) GetUser(ID string) (string, error) {
	return rb.IInfraPersistenceRedis.GetKey(ID)
}
