package update

type InputOrderUpdateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
}

type OutputOrderUpdateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
}
