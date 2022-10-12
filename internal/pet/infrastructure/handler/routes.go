package handler

import (
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, h Handler) {
	r.HandleFunc("/pet/{name}", h.Get).Methods("GET")
	r.HandleFunc("/pet", h.Add).Methods("POST")
}
