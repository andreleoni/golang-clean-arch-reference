package create

import (
	"testing"
)

func TestUseCaseCustomerCreateHandler_Integration(t *testing.T) {
	// postgres.PGSetup()

	// customerRepository := customer.NewCustomer(postgres.PG)

	// expectedName := "my name"

	// t.Run("when creating the handler", func(t *testing.T) {
	// 	icfd := InputCustomerCreateDTO{Name: expectedName}

	// 	customerFindHandler := NewUseCaseCustomerCreateHandler(customerRepository)

	// 	expectedResult := OutputCustomerCreateDTO{Name: expectedName}

	// 	result, err := customerFindHandler.Handle(icfd)

	// 	assert.Nil(t, err)
	// 	assert.NotEqual(t, result.ID, "")
	// 	assert.Equal(t, expectedResult.Name, result.Name)

	// 	findResult, err := customerRepository.Find(result.ID)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, result.ID, findResult.ID)
	// 	assert.Equal(t, expectedResult.Name, result.Name)
	// })
}
