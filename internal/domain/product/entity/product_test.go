package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testproduct_Validate(t *testing.T) {
	product := Product{}

	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("ID can't be blank,product can't be blank,customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.ID = "testing-uuid"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("product can't be blank,customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.Name = "testing-product-uuid"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("customer can't be blank,quantity must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.CustomerID = "testing-customer-uuid"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("quantity must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.Quantity = 1

	product.Validate()
	assert.Equal(t, product.HasErrors(), false)
}
