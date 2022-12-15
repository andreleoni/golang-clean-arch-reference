package entities

import "fmt"

type Notificator struct {
	Errors []error `sql:"-"`
}

func (c *Notificator) HasErrors() bool {
	return len(c.Errors) > 0
}

func (c *Notificator) ErrorMessage() error {
	fullMessage := ""

	for _, currentError := range c.Errors {
		fullMessage += currentError.Error()
		fullMessage += ","
	}

	return fmt.Errorf(fullMessage)
}
