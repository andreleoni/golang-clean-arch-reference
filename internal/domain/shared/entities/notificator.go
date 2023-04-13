package entities

import "fmt"

type Notificator struct {
	Errors []error `gorm:"-"`
}

func (n *Notificator) Validate() {
}

func (n *Notificator) AddError(err error) {
	n.Errors = append(n.Errors, err)
}

func (n *Notificator) HasErrors() bool {
	return len(n.Errors) > 0
}

func (n *Notificator) ErrorMessage() error {
	fullMessage := ""

	for _, currentError := range n.Errors {
		fullMessage += currentError.Error()
		fullMessage += ","
	}

	return fmt.Errorf(fullMessage)
}
