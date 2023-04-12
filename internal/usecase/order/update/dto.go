package update

import "golang-clean-arch-reference/internal/domain/valueobject"

type InputOrderUpdateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
	Address    valueobject.Address
}

type OutputOrderUpdateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
	Address    valueobject.Address
}
