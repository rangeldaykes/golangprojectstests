package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var fcvs []Fcv

// Fcv is a linha tabela
type Fcv struct {
	Codigo  string `json:"codigo"`
	IDLinha string `json:"idlinha"`
}

func main() {

	fcvs = append(fcvs, Fcv{Codigo: "1", IDLinha: "050"})

	r := mux.NewRouter()

	r.HandleFunc("/api/fcvs", getFcvs).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getFcvs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	n := []int{5, 7, 4}
	fmt.Println(n[3])

	json.NewEncoder(w).Encode(fcvs)
}
