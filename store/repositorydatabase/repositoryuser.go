package repositorydatabase

import (
	"store/domain/entities"
	"store/infra/databasesqlx"
)

type IRepositoryUser interface {
	GetUser(ID int) (entities.User, error)
}

type repositoryUser struct {
	databasesqlx.IInfraRepositorySqlx
}

func NewRepositoryUser(infra databasesqlx.IInfraRepositorySqlx) *repositoryUser {
	return &repositoryUser{infra}
}

func (ru repositoryUser) GetUser(ID int) (entities.User, error) {
	sql := GetUserSql()

	s := struct {
		ID int
	}{
		ID,
	}

	var ret entities.User
	err := ru.PrepareNamedGetOne(&ret, sql, s)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
