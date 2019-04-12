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
	ret, err := rs.IRepositoryCacheUser.GetUser(strconv.Itoa(iduser))
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
