package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"kittyhome/internal/pet/application"
	"kittyhome/internal/pet/domain"
	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase application.Usecase
}

func Build(u application.Usecase) Handler {
	return &handler{usecase: u}
}

// Get a pet by name
// @Summary     Get a pet by name
// @Description get a pet by name
// @Tags        pets
// @Produce     json
// @Param       name path     string true "Nymeria"
// @Success     200  {object} domain.Pet
// @Failure     400  {object} errors.CustomError
// @Failure     404  {object} errors.CustomError
// @Failure     500  {object} errors.CustomError
// @Router      /pet/{name} [get]
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	// name := req.URL.Query().Get("name")
	p, err := h.usecase.Get(name)
	if err != nil {
		log.Err(err)
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Err(err)
	}
}

// Add a pet
// @Summary     Add a pet
// @Description add a pet with autoincrement id
// @Tags        pets
// @Accept      json
// @Param       pet body domain.Pet true "Pet data"
// @Success     200
// @Failure     400 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /pet [post]
func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	p := &domain.Pet{}
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		log.Err(err)
	}
	_, err = h.usecase.Add(p)
	if err != nil {
		log.Err(err)
	}
	w.WriteHeader(http.StatusOK)
}
