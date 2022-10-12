package repository

import (
	"fmt"
	"kittyhome/internal/pet/domain"
)

type Mysql interface {
	Add(p *domain.Pet) (int, error)
	Get(name string) (*domain.Pet, error)
}

type mysql struct {
	pets   []*domain.Pet
	lastId int
}

func Build() Mysql {
	return &mysql{pets: []*domain.Pet{
		{
			Color:    "black",
			Type:     "cat",
			Pedigree: "normal",
			Name:     "Nymeria",
			Id:       1,
		}, {
			Color:    "carey",
			Type:     "cat",
			Pedigree: "carey",
			Name:     "Lyanna",
			Id:       2,
		},
	},
		lastId: 2,
	}
}

func (m *mysql) Add(p *domain.Pet) (int, error) {
	m.lastId++
	p.Id = m.lastId
	m.pets = append(m.pets, p)
	return p.Id, nil
}

func (m *mysql) Get(name string) (*domain.Pet, error) {
	for _, p := range m.pets {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("pet %s not found", name)
}
