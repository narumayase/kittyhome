package repository

import (
	"kittyhome/internal/toy/domain"
)

type Elastic interface {
	Add(t *domain.Toy) (int, error)
	Get() ([]*domain.Toy, error)
}

type elastic struct {
	toys   []*domain.Toy
	lastId int
}

func Build() Elastic {
	return &elastic{toys: []*domain.Toy{
		{
			Brand: "black",
			Owner: "Nymeria",
			Id:    1,
		}, {
			Brand: "white",
			Owner: "Lyanna",
			Id:    2,
		},
	},
		lastId: 2,
	}
}

func (e *elastic) Add(t *domain.Toy) (int, error) {
	e.lastId++
	t.Id = e.lastId
	e.toys = append(e.toys, t)
	return t.Id, nil
}

func (e *elastic) Get() ([]*domain.Toy, error) {
	return e.toys, nil
}
