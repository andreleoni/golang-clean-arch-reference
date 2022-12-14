package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer_Validate(t *testing.T) {
	customer := Customer{}

	idBlankError := customer.Validate()
	if idBlankError != nil {
		assert.Equal(t, fmt.Errorf("ID can't be blank"), idBlankError)
	}

	customer.ID = "testing-uuid"
	nameBlankError := customer.Validate()
	if customer.Validate() != nil {
		assert.Equal(t, fmt.Errorf("name can't be blank"), nameBlankError)
	}

	customer.Name = "Customer Name"

	assert.Nil(t, customer.Validate())
}
