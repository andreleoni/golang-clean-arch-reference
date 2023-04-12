package find

import "golang-clean-arch-reference/internal/domain/valueobject"

type InputOrderFindDTO struct {
	ID string
}

type OutputOrderFindDTO struct {
	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
	Address    valueobject.Address
}
