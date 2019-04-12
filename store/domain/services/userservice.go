package services

import (
	"encoding/json"
	"store/domain/interfaces"
	"store/domain/models/viewmodels"
	"strconv"
)

type userService struct {
	interfaces.IRepositoryCacheUser
	interfaces.IRepositoryDispMobile
}

func NewUserService(
	cache interfaces.IRepositoryCacheUser,
	repo interfaces.IRepositoryDispMobile) *userService {
	return &userService{cache, repo}
}

func (rs userService) GetUser(iduser int) (viewmodels.UserVM, error) {
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
