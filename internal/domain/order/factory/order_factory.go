package factory

import (
	"golang-clean-arch-reference/internal/domain/order/entity"

	"github.com/google/uuid"
)

type OrderFactory struct{}

func NewOrderFactory() *OrderFactory {
	return &OrderFactory{}
}

func (cf *OrderFactory) Create(id, productID, customerID string, quantity uint64) *entity.Order {
	if id == "" {
		id = uuid.New().String()
	}

	order := entity.Order{
		ID:         id,
		ProductID:  productID,
		CustomerID: customerID,
		Quantity:   quantity,
	}

	return &order
}
