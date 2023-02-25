package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Validate(t *testing.T) {
	product := Product{}

	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("ID can't be blank,name can't be blank,status can't be blank,price must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.Status = "enabled"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("ID can't be blank,name can't be blank,price must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.ID = "my-product-id"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("name can't be blank,price must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.Name = "product name"
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf("price must be great than zero,"))
	assert.Equal(t, product.HasErrors(), true)

	product.Price = 100
	product.Validate()
	assert.Equal(t, product.ErrorMessage(), fmt.Errorf(""))
	assert.Equal(t, product.HasErrors(), false)
}
