package factory

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"

	"github.com/google/uuid"
)

func NewCustomer(id, name string) entity.Customer {
	if id == "" {
		id = uuid.New().String()
	}

	return entity.Customer{
		ID:   id,
		Name: name,
	}
}
