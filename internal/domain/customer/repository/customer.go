package repository

import "golang-clean-arch-reference/internal/domain/customer/entity"

type Customer interface {
	Get(ID string) (*entity.Customer, error)
	Create(*entity.Customer) error
	Update(*entity.Customer) error
	Delete(*entity.Customer) error
}
