package interfaces

import (
	"store/domain/models/entities"
)

type IRepositoryCacheUser interface {
	GetUser(ID string) (entities.User, error)
}
