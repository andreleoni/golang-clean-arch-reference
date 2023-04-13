package create

import "golang-clean-arch-reference/internal/domain/valueobject"

type InputOrderCreateDTO struct {
	ProductID  string
	CustomerID string
	Quantity   uint64
	Address    valueobject.Address
}

type OutputOrderCreateDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
	Address    valueobject.Address
}
