package controllers

import (
	"net/http"
)

func ErrorHandling(w http.ResponseWriter, err error) {
	if err != nil {

		//log.Println(string(debug.Stack()))
		//http.Error(w, fmt.Sprintf("%s \n %s", err.Error(), string(debug.Stack())), 500)

		http.Error(w, err.Error(), 500)
	}
}
