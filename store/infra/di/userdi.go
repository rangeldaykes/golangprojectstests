package di

import (
	"store/api/controllers"
	"store/domain/services/userservice"
	"store/infra/cacheredis/redisredigo"
	"store/repositorycache"
	"store/repositorydatabase"
)

func InjectUserController() *controllers.UserController {
	databasesqlx := injectdatabase()
	repodatabasedbdisp := repositorydatabase.NewRepositoryUser(databasesqlx)
	repocache := redisredigo.NewPersistenceRedigo()
	repocacheuser := repositorycache.NewRepositoryCacheUser(repocache)
	serviceuser := userservice.NewUserService(repocacheuser, repodatabasedbdisp)
	controller := controllers.NewUserController(serviceuser)

	return controller
}
