package userservice

import (
	"store/domain/entities"
	"store/domain/viewmodels"
	"store/repositorycache"
	"store/repositorydatabase/repositoryuser"
	"strconv"
)

// IUserService for user methods
type IUserService interface {
	GetUser(iduser int) (viewmodels.UserVM, error)
}

type userService struct {
	repositorycache.IRepositoryCacheUser
	repositoryuser.IRepositoryUser
}

func NewUserService(
	cache repositorycache.IRepositoryCacheUser,
	repo repositoryuser.IRepositoryUser) *userService {
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
		user, err = rs.IRepositoryUser.GetUser(iduser)
		if err != nil {
			return viewmodels.UserVM{}, err
		}
	}

	uservm := viewmodels.UserVM{ID: user.ID, Name: user.Name}
	return uservm, nil
}
