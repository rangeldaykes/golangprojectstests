package di

import "store/infra/databasesqlx"

func injectdatabase() *databasesqlx.InfraRepositorySqlx {
	repoconnsqlx := databasesqlx.ConnDataBaseGetInstance(true)
	databasesqlx := databasesqlx.NewInfraRepositorySqlx(repoconnsqlx)
	return databasesqlx
}
