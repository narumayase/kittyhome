package application

import (
	"kittyhome/internal/toy/domain"
)

type Usecase interface {
	Add(p *domain.Toy) (int, error)
	Get() ([]*domain.Toy, error)
}
