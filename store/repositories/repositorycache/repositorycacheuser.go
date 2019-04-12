package repositorycache

import (
	"encoding/json"
	"store/domain/models/entities"
	"store/infrastucture/data/infraredis"
)

type repositoryCacheUser struct {
	infraredis.IInfraPersistenceRedis
}

func NewRepositoryCacheUser(infra infraredis.IInfraPersistenceRedis) *repositoryCacheUser {
	return &repositoryCacheUser{infra}
}

func (rb repositoryCacheUser) GetUser(ID string) (entities.User, error) {
	ret, err := rb.IInfraPersistenceRedis.GetKey(ID)
	if err != nil || ret == "" {
		return entities.User{}, err
	}

	user := entities.User{}
	err = json.Unmarshal([]byte(ret), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
