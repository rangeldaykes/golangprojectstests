package controllers

import (
	"encoding/json"
	"net/http"
	"service-pattern-go-master/interfaces"
	"service-pattern-go-master/viewmodels"

	"github.com/go-chi/chi"
)

type PlayerController struct {
	interfaces.IPlayerService
}

func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{scores})
}
