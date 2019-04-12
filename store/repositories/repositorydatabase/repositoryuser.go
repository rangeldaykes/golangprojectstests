package repositorydatabase

import (
	"store/domain/models/entities"
	"store/infrastucture/data/infradatabasesqlx"
)

type repositoryUser struct {
	infradatabasesqlx.IInfraRepositorySqlx
}

func NewRepositoryUser(infra infradatabasesqlx.IInfraRepositorySqlx) *repositoryUser {
	return &repositoryUser{infra}
}

func (ru repositoryUser) GetUser(ID int) (entities.User, error) {
	sql := `
	select u.id, u.name
      from users u
     where u.id = :id`

	s := struct {
		ID int
	}{
		ID,
	}

	var ret entities.User
	err := ru.PrepareNamedGetOne(&ret, sql, s)
	if err != nil {
		// log err
		return ret, nil
	}

	return ret, nil
}
