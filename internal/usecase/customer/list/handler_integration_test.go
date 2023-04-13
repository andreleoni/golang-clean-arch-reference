package list

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseCustomerListHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	customerRepository := customer.NewCustomer(sqlite.Sqlite)

	firstItemID := "my-uuid-1"
	firstItemName := "my name 1"

	customer1 := entity.Customer{ID: firstItemID, Name: firstItemName}

	err := customerRepository.Create(&customer1)
	assert.Nil(t, err)

	secondItemID := "my-uuid-2"
	secondItemName := "my name 1"

	customer2 := entity.Customer{ID: secondItemID, Name: secondItemName}

	err = customerRepository.Create(&customer2)
	assert.Nil(t, err)

	t.Run("when have records on database", func(t *testing.T) {
		icfd := InputCustomerListDTO{}

		customerListHandler := NewUseCaseCustomerListHandler(customerRepository)

		result, err := customerListHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, len(result.Customers), 2)

		assert.Equal(t, result.Customers[0].ID, firstItemID)
		assert.Equal(t, result.Customers[0].Name, firstItemName)

		assert.Equal(t, result.Customers[1].ID, secondItemID)
		assert.Equal(t, result.Customers[1].Name, secondItemName)
	})
}
