package find

import (
	"testing"
)

func TestUseCaseCustomerFindHandler_Integration(t *testing.T) {
	// postgres.PGSetup()

	// customerRepository := customer.NewCustomer(postgres.PG)

	// expectedID := "my-uuid"
	// expectedName := "my name"

	// customer := entity.Customer{ID: expectedID, Name: expectedName}

	// err := customerRepository.Create(&customer)
	// assert.Nil(t, err)

	// t.Run("when have the ID on database", func(t *testing.T) {
	// 	icfd := InputCustomerFindDTO{ID: expectedID}

	// 	customerFindHandler := NewUseCaseCustomerFindHandler(customerRepository)

	// 	expectedResult := OutputCustomerFindDTO{ID: expectedID, Name: expectedName}

	// 	result, err := customerFindHandler.Handle(icfd)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, expectedResult, result)
	// })

	// t.Run("when have not the ID on database", func(t *testing.T) {
	// 	icfd := InputCustomerFindDTO{ID: "not existing ID"}

	// 	customerFindHan	dler := NewUseCaseCustomerFindHandler(customerRepository)

	// 	expectedResult := OutputCustomerFindDTO{}

	// 	result, err := customerFindHandler.Handle(icfd)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, expectedResult, result)
	// })
}
