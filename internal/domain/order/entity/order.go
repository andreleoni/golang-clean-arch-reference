package entity

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/shared/entities"
)

type Order struct {
	entities.Notificator

	ID         string
	ProductID  string
	CustomerID string
	Quantity   uint64
}

func (c *Order) Validate() {
	c.Errors = []error{}

	if c.ID == "" {
		c.AddError(fmt.Errorf("ID can't be blank"))
	}

	if c.ProductID == "" {
		c.AddError(fmt.Errorf("product can't be blank"))
	}

	if c.CustomerID == "" {
		c.AddError(fmt.Errorf("customer can't be blank"))
	}

	if c.Quantity <= 0.0 {
		c.AddError(fmt.Errorf("quantity must be great than zero"))
	}
}
