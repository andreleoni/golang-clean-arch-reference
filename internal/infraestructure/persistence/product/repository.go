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

func (c Product) Create(ec *entity.Product) error {
	return c.pg.Insert(ec)
}

func (c Product) Update(ec *entity.Product) error {
	_, err := c.pg.Model(ec).Set("name = ?name, status = ?status, price = ?price").Where("id = ?id").Update()

	return err
}

func (c Product) Delete(ec *entity.Product) error {
	_, err := c.pg.Model(ec).Where("id = ?id").Delete()

	return err
}
