package interfaces

import (
	"store/domain/models/viewmodels"
)

// IUserService for user methods
type IUserService interface {
	GetUser(iduser int) (viewmodels.UserVM, error)
}
