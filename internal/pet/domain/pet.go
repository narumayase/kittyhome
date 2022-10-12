package domain

// Pet
// @Description
type Pet struct {
	// Color of the pet
	Color string
	// Type of the pet (cat, dog)
	Type string
	// Pedigree of the pet
	Pedigree string
	// Name of the pet
	Name string
	// Id of the pet
	Id int
}

func (c *Pet) GetPedigree() string {
	return c.Pedigree
}

type PetRepository interface {
	Add(*Pet) (int, error)
	Get(name string) (*Pet, error)
}
