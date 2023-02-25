package customer

import (
	"golang-clean-arch-reference/internal/domain/product/entity"

	"github.com/go-pg/pg"
)

type Product struct {
	pg *pg.DB
}

func NewProduct(pg *pg.DB) Product {
	return Product{pg: pg}
}

func (c Product) List() ([]*entity.Product, error) {
	product := []*entity.Product{}
	err := c.pg.Model(&product).Select()

	return product, err
}

func (c Product) Find(ID string) (*entity.Product, error) {
	product := entity.Product{ID: ID}
	err := c.pg.Model(&product).WherePK().Select()

	return &product, err
}

func (c Product) Create(ep *entity.Product) error {
	return c.pg.Insert(ep)
}

func (c Product) Update(ep *entity.Product) error {
	_, err := c.pg.Model(ep).Set("name = ?name, status = ?status, price = ?price").Where("id = ?id").Update()

	return err
}

func (c Product) Enable(ep *entity.Product) error {
	ep.Status = "enabled"

	return c.Update(ep)
}

func (c Product) Disable(ep *entity.Product) error {
	ep.Status = "disabled"

	return c.Update(ep)
}
