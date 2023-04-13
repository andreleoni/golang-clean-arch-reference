package find

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseCustomerFindHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	customerRepository := customer.NewCustomer(sqlite.Sqlite)

	expectedID := "my-uuid"
	expectedName := "my name"

	customer := entity.Customer{ID: expectedID, Name: expectedName}

	err := customerRepository.Create(&customer)
	assert.Nil(t, err)

	t.Run("when have the ID on database", func(t *testing.T) {
		icfd := InputCustomerFindDTO{ID: expectedID}

		customerFindHandler := NewUseCaseCustomerFindHandler(customerRepository)

		expectedResult := OutputCustomerFindDTO{ID: expectedID, Name: expectedName}

		result, err := customerFindHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("when have not the ID on database", func(t *testing.T) {
		icfd := InputCustomerFindDTO{ID: "not existing ID"}

		customerFindHandler := NewUseCaseCustomerFindHandler(customerRepository)

		expectedResult := OutputCustomerFindDTO{}

		result, err := customerFindHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
