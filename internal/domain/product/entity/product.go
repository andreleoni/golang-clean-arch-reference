package entity

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/shared/entities"
)

type Product struct {
	entities.Notificator

	ID     string
	Name   string
	Status string
	Price  uint64
}

func (c *Product) Validate() {
	c.Errors = []error{}

	if c.ID == "" {
		c.AddError(fmt.Errorf("ID can't be blank"))
	}

	if c.Name == "" {
		c.AddError(fmt.Errorf("name can't be blank"))
	}

	if c.Status == "" {
		c.AddError(fmt.Errorf("status can't be blank"))
	}

	if c.Price <= 0.0 {
		c.AddError(fmt.Errorf("price must be great than zero"))
	}
}
