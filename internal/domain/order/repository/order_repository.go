package repository

import "golang-clean-arch-reference/internal/domain/order/entity"

type OrderRepository interface {
	Find(ID string) (*entity.Order, error)
	Create(*entity.Order) error
	Update(*entity.Order) error
}
