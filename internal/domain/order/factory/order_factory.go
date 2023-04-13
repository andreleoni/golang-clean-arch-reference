package factory

import (
	"golang-clean-arch-reference/internal/domain/order/entity"
	"golang-clean-arch-reference/internal/domain/valueobject"

	"github.com/google/uuid"
)

type OrderFactory struct{}

func NewOrderFactory() *OrderFactory {
	return &OrderFactory{}
}

func (cf *OrderFactory) Create(productID, customerID string, quantity uint64, address valueobject.Address) *entity.Order {
	id := uuid.New().String()

	order := entity.Order{
		ID:         id,
		ProductID:  productID,
		CustomerID: customerID,
		Quantity:   quantity,
		Address:    address,
	}

	return &order
}
