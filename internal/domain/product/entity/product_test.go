package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_Validate(t *testing.T) {
	order := Order{}

	order.Validate()
	assert.Equal(t, order.ErrorMessage(), fmt.Errorf("ID can't be blank,product can't be blank,customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, order.HasErrors(), true)

	order.ID = "testing-uuid"
	order.Validate()
	assert.Equal(t, order.ErrorMessage(), fmt.Errorf("product can't be blank,customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, order.HasErrors(), true)

	order.ProductID = "testing-product-uuid"
	order.Validate()
	assert.Equal(t, order.ErrorMessage(), fmt.Errorf("customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, order.HasErrors(), true)

	order.CustomerID = "testing-customer-uuid"
	order.Validate()
	assert.Equal(t, order.ErrorMessage(), fmt.Errorf("quantity must be great than zero,"))
	assert.Equal(t, order.HasErrors(), true)

	order.Quantity = 1

	order.Validate()
	assert.Equal(t, order.HasErrors(), false)
}
