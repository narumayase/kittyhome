package application

import (
	"kittyhome/internal/pet/domain"
)

type Usecase interface {
	Add(p *domain.Pet) (int, error)
	Get(name string) (*domain.Pet, error)
}
