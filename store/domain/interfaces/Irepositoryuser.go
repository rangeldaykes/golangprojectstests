package interfaces

import (
	"store/domain/models/entities"
)

// IRepositoryDispMobile interface for dispositives mobiles
type IRepositoryDispMobile interface {
	GetUser(ID int) (entities.User, error)
}
