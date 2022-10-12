package domain

// Toy
// @Description toy data
type Toy struct {
	// Brand of the toy
	Brand string
	// Owner pet owner of the toy
	Owner string
	// Id of the toy
	Id int
}

// ToyRepository interface
type ToyRepository interface {
	Add(*Toy) (int, error)
	Get() ([]*Toy, error)
}
