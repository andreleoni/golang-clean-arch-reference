package customer

import (
	"golang-clean-arch-reference/internal/domain/order/entity"

	"gorm.io/gorm"
)

type Order struct {
	sqlite *gorm.DB
}

func NewOrder(sqlite *gorm.DB) Order {
	return Order{sqlite: sqlite}
}

func (c Order) Find(ID string) (*entity.Order, error) {
	order := entity.Order{ID: ID}
	result := c.sqlite.Find(&order)

	return &order, result.Error
}

func (c Order) Create(o *entity.Order) error {
	return c.sqlite.Create(&o).Error
}

func (c Order) Update(o *entity.Order) error {
	return c.sqlite.Save(&o).Error
}
