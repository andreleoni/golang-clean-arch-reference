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

func (c Customer) Find(ID string) (*entity.Customer, error) {
	return &entity.Customer{}, nil
}

func (c Customer) Create(*entity.Customer) error {
	log.Fatal("not implemented error")

	return nil
}

func (c Customer) Update(*entity.Customer) error {
	log.Fatal("not implemented error")

	return nil
}

func (c Customer) Delete(*entity.Customer) error {
	log.Fatal("not implemented error")

	return nil
}
