package repositorydatabase

import "store/infrastucture/data/infradatabasesqlx"

type repositoryUser struct {
	infradatabasesqlx.IInfraRepositorySqlx
}

func NewRepositoryUser(infra infradatabasesqlx.IInfraRepositorySqlx) *repositoryUser {
	return &repositoryUser{infra}
}

func (rd repositoryUser) CarregarGoogleId(codigoVeiculo int) string {

	sql := `
	select c.name
      from clients c
     where c.id = :id`

	s := struct {
		ID int
	}{
		codigoVeiculo,
	}

	var ret string
	rd.PrepareNamedGetOne(&ret, sql, s)

	return ret
}
