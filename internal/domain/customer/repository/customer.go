package repository

import "golang-clean-arch-reference/internal/domain/customer/entity"

type Customer interface {
	List() ([]*entity.Customer, error)
	Find(ID string) (*entity.Customer, error)
	Create(*entity.Customer) error
	Update(*entity.Customer) error
	Delete(*entity.Customer) error
}
