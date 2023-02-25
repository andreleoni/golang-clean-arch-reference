package find

type InputProductFindDTO struct {
	ID string
}

type OutputProductFindDTO struct {
	ID     string
	Name   string
	Status string
	Price  uint64
}
