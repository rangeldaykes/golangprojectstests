package injectioncontainer

import (
	"store/api/controllers"
	"store/domain/services"
	"store/infrastucture/data/infradatabasesqlx"
	"store/infrastucture/data/infraredis/infraredigo"
	"store/repositories/repositorycache"
	"store/repositories/repositorydatabase"
)

func InjectUserController() *controllers.UserController {
	databasesqlx := injectdatabase()
	repodatabasedbdisp := &repositorydatabase.RepositoryDispMobile{IInfraRepositorySqlx: databasesqlx}
	repocache := &infraredigo.PersistenceRedigo{}
	repocacheuser := &repositorycache.RepositoryCacheUser{IInfraPersistenceRedis: repocache}
	service := &services.UserService{IRepositoryCacheUser: repocacheuser, IRepositoryDispMobile: repodatabasedbdisp}
	controller := controllers.UserController{IUserService: service}
	return &controller
}

func injectdatabase() *infradatabasesqlx.InfraRepositorySqlx {
	repoconnsqlx := infradatabasesqlx.ConnDataBaseGetInstance(true)
	databasesqlx := &infradatabasesqlx.InfraRepositorySqlx{IConnDataBase: repoconnsqlx}

	return databasesqlx
}
