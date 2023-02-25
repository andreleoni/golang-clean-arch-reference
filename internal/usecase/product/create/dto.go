package create

type InputProductCreateDTO struct {
	Name   string
	Status string
	Price  uint64
}

type OutputProductCreateDTO struct {
	ID     string
	Name   string
	Status string
	Price  uint64
}
