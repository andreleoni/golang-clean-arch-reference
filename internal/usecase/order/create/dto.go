package create

type InputOrderCreateDTO struct {
	ProductID  string
	CustomerID string
	Quantity   uint64
}

type OutputOrderCreateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
}
