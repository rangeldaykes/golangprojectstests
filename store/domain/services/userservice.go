package services

import (
	"encoding/json"
	"store/domain/interfaces"
	"store/domain/models/viewmodels"
	"strconv"
)

// UserService implements a services for user
type UserService struct {
	interfaces.IRepositoryCacheUser
	interfaces.IRepositoryDispMobile
}

func (rs UserService) GetUser(iduser int) (viewmodels.UserVM, error) {
	ret, err := rs.GetKey(strconv.Itoa(iduser))
	if err != nil || ret == "" {
		return viewmodels.UserVM{}, err
	}

	user := viewmodels.UserVM{}
	err = json.Unmarshal([]byte(ret), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// ProcessarPedidoFcv return fcv from cache
func (rs UserService) processarPedidoFcv(id string, fcvs []int) (viewmodels.UserVM, error) {
	//return valueobject.FCVPlanejamentoVO{IDVeiculo: idVeiculo}, nil

	ret, err := rs.GetKey(id)
	if err != nil {
		return viewmodels.UserVM{}, err
	}

	user := viewmodels.UserVM{}
	err = json.Unmarshal([]byte(ret), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (rs UserService) CarregarGoogleId(codigoVeiculo int) string {
	return rs.IRepositoryDispMobile.CarregarGoogleId(codigoVeiculo)
}
