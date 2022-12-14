package factory

import "golang-clean-arch-reference/internal/domain/customer/entity"

func NewCustomer(id, name string) entity.Customer {
	return entity.Customer{
		ID:   id,
		Name: name,
	}
}
