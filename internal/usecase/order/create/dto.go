package create

type InputOrderFindDTO struct {
	ID string
}

type OutputOrderFindDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
}
