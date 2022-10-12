package handler

import (
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router, h Handler) {
	r.HandleFunc("/toy", h.Get).Methods("GET")
	r.HandleFunc("/toy", h.Add).Methods("POST")
}
