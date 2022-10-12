package controller

import (
	"kittyhome/internal/toy/application"
	"kittyhome/internal/toy/domain"
)

type controller struct {
	repository domain.ToyRepository
}

func Build(r domain.ToyRepository) application.Usecase {
	return &controller{repository: r}
}

func (c *controller) Add(p *domain.Toy) (int, error) {
	return c.repository.Add(p)
}

func (c *controller) Get() ([]*domain.Toy, error) {
	return c.repository.Get()
}
