package controllers

import (
	"net/http"
)

func ErrorHandling(w http.ResponseWriter, err error) {
	if err != nil {
		//log.Printf("error in List - error: %v", err)
		//w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), 500)
	}
}
