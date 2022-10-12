package handler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"kittyhome/internal/toy/application"
	"kittyhome/internal/toy/domain"
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

// Get all toys
// @Summary     Get all toys
// @Description get all toys
// @Tags        toys
// @Produce     json
// @Success     200 {object} domain.Toy
// @Failure     400 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /toy [get]
func (h *handler) Get(w http.ResponseWriter, _ *http.Request) {
	p, err := h.usecase.Get()
	if err != nil {
		log.Err(err)
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Err(err)
	}
}

// Add a toy
// @Summary     Add a toy
// @Description add a toy with autoincrement id
// @Tags        toys
// @Accept      json
// @Param       toy body domain.Toy true "Toy data"
// @Success     200
// @Failure     400 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /toy [post]
func (h *handler) Add(w http.ResponseWriter, r *http.Request) {
	p := &domain.Toy{}
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
