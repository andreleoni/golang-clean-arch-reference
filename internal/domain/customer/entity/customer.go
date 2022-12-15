package entity

import "fmt"

type Customer struct {
	ID   string
	Name string

	errors []error
}

func (c *Customer) Validate() {
	c.errors = []error{}

	if c.ID == "" {
		c.errors = append(c.errors, fmt.Errorf("ID can't be blank"))
	}

	if c.Name == "" {
		c.errors = append(c.errors, fmt.Errorf("name can't be blank"))
	}
}

func (c *Customer) HasErrors() bool {
	return len(c.errors) > 0
}

func (c *Customer) ErrorMessage() error {
	fullMessage := ""
	for _, currentCustomer := range c.errors {
		fullMessage += currentCustomer.Error()
		fullMessage += ","
	}

	return fmt.Errorf(fullMessage)
}
