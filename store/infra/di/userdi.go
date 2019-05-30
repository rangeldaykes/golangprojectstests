package di

import (
	"store/api/controllers"
	"store/domain/services/userservice"
	"store/infra/cacheredis/redisredigo"
	"store/repositorycache"
	"store/repositorydatabase/repositoryuser"
)

func InjectUserController() *controllers.UserController {
	databasesqlx := injectdatabase()
	repodatabasedbdisp := repositoryuser.NewRepositoryUser(databasesqlx)
	repocache := redisredigo.NewPersistenceRedigo()
	repocacheuser := repositorycache.NewRepositoryCacheUser(repocache)
	serviceuser := userservice.NewUserService(repocacheuser, repodatabasedbdisp)
	controller := controllers.NewUserController(serviceuser)

	return controller
}
