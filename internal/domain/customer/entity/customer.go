package entity

import (
	"fmt"
)

type Customer struct {
	ID   string
	Name string
}

func (c Customer) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("ID can't be blank")
	}

	if c.Name == "" {
		return fmt.Errorf("name can't be blank")
	}

	return nil
}
