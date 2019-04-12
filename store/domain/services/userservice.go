package services

import (
	"store/domain/interfaces"
	"store/domain/models/entities"
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
	var user = entities.User{}
	var err error
	user, err = rs.IRepositoryCacheUser.GetUser(strconv.Itoa(iduser))
	if err != nil {
		//return viewmodels.UserVM{}, err
		// log
	}

	if user.ID == 0 {
		user, err = rs.IRepositoryDispMobile.GetUser(iduser)
		if err != nil {
			return viewmodels.UserVM{}, err
		}
	}

	uservm := viewmodels.UserVM{ID: user.ID, Name: user.Name}
	return uservm, nil
}
