package delete

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseCustomerDeleteHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	customerRepository := customer.NewCustomer(sqlite.Sqlite)

	expectedID := "my-uuid"
	expectedName := "my name"

	customer := entity.Customer{ID: expectedID, Name: expectedName}

	err := customerRepository.Create(&customer)
	assert.Nil(t, err)

	t.Run("when deleting the record", func(t *testing.T) {
		customerDeleteHandler := NewUseCaseCustomerDeleteHandler(customerRepository)

		icfd := InputCustomerDeleteDTO{ID: expectedID}

		_, err := customerDeleteHandler.Handle(icfd)
		assert.Nil(t, err)

		_, err = customerRepository.Find(expectedID)

		assert.Equal(t, err.Error(), "record not found")
	})
}
