package controller

import (
	"kittyhome/internal/pet/application"
	"kittyhome/internal/pet/domain"
)

type controller struct {
	repository domain.PetRepository
}

func Build(r domain.PetRepository) application.Usecase {
	return &controller{repository: r}
}

func (u *controller) Add(p *domain.Pet) (int, error) {
	return u.repository.Add(p)
}

func (u *controller) Get(name string) (*domain.Pet, error) {
	return u.repository.Get(name)
}
