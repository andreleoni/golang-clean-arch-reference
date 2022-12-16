package create

type InputProductCreateDTO struct {
	Name  string
	Price uint64
}

type OutputProductCreateDTO struct {
	ID    string
	Name  string
	Price uint64
}
