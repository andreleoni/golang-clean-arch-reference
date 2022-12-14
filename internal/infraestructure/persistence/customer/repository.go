package customer

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"
	"log"

	"github.com/go-pg/pg"
)

type Customer struct {
	pg *pg.DB
}

func NewCustomer(pg *pg.DB) Customer {
	return Customer{pg: pg}
}

func (c Customer) List() ([]*entity.Customer, error) {
	customer := []*entity.Customer{}
	err := c.pg.Model(&customer).Select()

	return customer, err
}

func (c Customer) Find(ID string) (*entity.Customer, error) {
	customer := entity.Customer{ID: ID}
	err := c.pg.Model(&customer).Select()

	return &customer, err
}

func (c Customer) Create(ec *entity.Customer) error {
	return c.pg.Insert(ec)
}

func (c Customer) Update(*entity.Customer) error {
	log.Fatal("not implemented error")

	return nil
}

func (c Customer) Delete(*entity.Customer) error {
	log.Fatal("not implemented error")

	return nil
}
