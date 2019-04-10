package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"store/domain/interfaces"
	"strconv"

	"github.com/go-chi/chi"
)

type UserController struct {
	interfaces.IUserService
}

// func ErrorHandling(w http.ResponseWriter, err error) {
// 	if err != nil {
// 		log.Printf("error in List - error: %v", err)
// 		//w.WriteHeader(http.StatusInternalServerError)
// 		http.Error(w, err.Error(), 500)
// 	}
// }

func (rs *UserController) Get(w http.ResponseWriter, r *http.Request) {

	userid, err := strconv.Atoi(chi.URLParam(r, "userid4"))
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

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ErrorHandling(w, err)
		return
	}
}

func (rs *UserController) List(w http.ResponseWriter, r *http.Request) {

	var err error
	defer func() {
		if err != nil {
			log.Printf("error in List - error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	userid, err := strconv.Atoi(chi.URLParam(r, "userid"))
	///var fcvs []int

	//fcv, err = rs.ProcessarPedidoFcv(veiculoid, fcvs)

	user, err := rs.GetUser(userid)

	//w.Write([]byte(idveiculo))

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(user)
}
