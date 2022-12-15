package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer_Validate(t *testing.T) {
	customer := Customer{}

	customer.Validate()
	assert.Equal(t, customer.ErrorMessage(), fmt.Errorf("ID can't be blank,name can't be blank,"))
	assert.Equal(t, customer.HasErrors(), true)

	customer.ID = "testing-uuid"
	customer.Validate()
	assert.Equal(t, customer.ErrorMessage(), fmt.Errorf("name can't be blank,"))
	assert.Equal(t, customer.HasErrors(), true)

	customer.Name = "Customer Name"

	customer.Validate()
	assert.Equal(t, customer.HasErrors(), false)
}
