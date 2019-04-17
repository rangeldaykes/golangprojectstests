package controllers

import (
	"net/http"
	"store/domain/interfaces"
	"strconv"

	"github.com/go-chi/chi"
)

type UserController struct {
	interfaces.IUserService
}

func NewUserController(userserv interfaces.IUserService) *UserController {
	return &UserController{userserv}
}

func (rs *UserController) Get(w http.ResponseWriter, r *http.Request) {
	userid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ErrorHandling(w, err)
		return
	}

	user, err := rs.GetUser(userid)
	if err != nil {
		ErrorHandling(w, err)
		return
	}

	if user.ID == 0 {
		http.Error(w, http.StatusText(404), 404)
	}

	//err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ErrorHandling(w, err)
		return
	}
}
