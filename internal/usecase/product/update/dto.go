package update

type InputProductUpdateDTO struct {
	ID     string
	Name   string
	Status string
	Price  uint64
}

type OutputProductUpdateDTO struct {
	ID     string
	Name   string
	Status string
	Price  uint64
}
