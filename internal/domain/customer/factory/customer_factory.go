package factory

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"

	"github.com/google/uuid"
)

type CustomerFactory struct{}

func NewCustomerFactory() *CustomerFactory {
	return &CustomerFactory{}
}

func (cf *CustomerFactory) Create(id, name string) *entity.Customer {
	if id == "" {
		id = uuid.New().String()
	}

	customer := entity.Customer{ID: id, Name: name}

	return &customer
}
