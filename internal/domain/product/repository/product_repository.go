package repository

import "golang-clean-arch-reference/internal/domain/product/entity"

type ProductRepository interface {
	List() ([]*entity.Product, error)
	Find(ID string) (*entity.Product, error)
	Create(*entity.Product) error
	Update(*entity.Product) error
	Enable(*entity.Product) error
	Disable(*entity.Product) error
}
