package customer

import (
	"golang-clean-arch-reference/internal/domain/order/entity"

	"github.com/go-pg/pg"
)

type Order struct {
	pg *pg.DB
}

func NewOrder(pg *pg.DB) Order {
	return Order{pg: pg}
}

func (c Order) Find(ID string) (*entity.Order, error) {
	order := entity.Order{ID: ID}
	err := c.pg.Model(&order).WherePK().Select()

	return &order, err
}

func (c Order) Create(ec *entity.Order) error {
	return c.pg.Insert(ec)
}

func (c Order) Update(ec *entity.Order) error {
	_, err := c.pg.Model(ec).Set(
		"product_id = ?product_id, customer_id = ?customer_id, quantity = ?quantity",
	).Where("id = ?id").Update()

	return err
}
