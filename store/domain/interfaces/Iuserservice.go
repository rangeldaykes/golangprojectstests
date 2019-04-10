package interfaces

import (
	"store/domain/models/viewmodels"
)

// IUserService for user methods
type IUserService interface {
	//ProcessarPedidoFcv(idVeiculo string, fcvs []int) (viewmodels.UserVM, error)
	GetUser(iduser int) (viewmodels.UserVM, error)
}
