package factory

import (
	"golang-clean-arch-reference/internal/domain/product/entity"

	"github.com/google/uuid"
)

type ProductFactory struct{}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (cf *ProductFactory) Create(name string, status string, price uint64) *entity.Product {
	id := uuid.New().String()

	product := entity.Product{
		ID:     id,
		Name:   name,
		Status: status,
		Price:  price,
	}

	return &product
}
