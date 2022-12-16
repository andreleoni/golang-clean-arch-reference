package update

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseCustomerUpdateHandler_Integration(t *testing.T) {
	postgres.PGSetup()

	customerRepository := customer.NewCustomer(postgres.PG)

	expectedID := "my-uuid"
	expectedName := "my name"

	customer := entity.Customer{ID: expectedID, Name: expectedName}

	err := customerRepository.Create(&customer)
	assert.Nil(t, err)

	t.Run("when updating the record", func(t *testing.T) {
		customerUpdateHandler := NewUseCaseCustomerUpdateHandler(customerRepository)

		myUpdatableName := "my another name"

		icfd := InputCustomerUpdateDTO{ID: expectedID, Name: myUpdatableName}

		result, err := customerUpdateHandler.Handle(icfd)
		assert.Nil(t, err)

		findResult, err := customerRepository.Find(result.ID)
		assert.Nil(t, err)

		assert.Equal(t, result.ID, findResult.ID)
		assert.Equal(t, result.Name, myUpdatableName)
	})
}
