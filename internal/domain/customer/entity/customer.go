package entity

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/shared/entities"
)

type Customer struct {
	entities.Notificator

	ID   string
	Name string
}

func (c *Customer) Validate() {
	c.Errors = []error{}

	if c.ID == "" {
		c.AddError(fmt.Errorf("ID can't be blank"))
	}

	if c.Name == "" {
		c.AddError(fmt.Errorf("name can't be blank"))
	}
}
