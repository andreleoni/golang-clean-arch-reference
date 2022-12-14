package find

import (
	"testing"
)

func TestUseCaseCustomerFindHandler_Handler(t *testing.T) {
	// t.Run("when the repository returns no errors", func(t *testing.T) {
	// 	customerMock := repository.MockCustomer{}

	// 	testingUUID := "my-uuid"
	// 	testingName := "My name"

	// 	customer := entity.Customer{ID: testingUUID, Name: testingName}

	// 	customerMock.On("Find", testingUUID).Return(&customer, nil)

	// 	uccfh := UseCaseCustomerFindHandler{customerRepository: &customerMock}

	// 	icfd := InputCustomerFindDTO{ID: testingUUID}

	// 	result, err := uccfh.Handle(icfd)

	// 	assert.Nil(t, err)
	// 	assert.Equal(t, result.ID, testingUUID)
	// 	assert.Equal(t, result.Name, testingName)
	// })

	// t.Run("when the repository returns errors", func(t *testing.T) {
	// 	customerMock := repository.MockCustomer{}

	// 	testingUUID := "my-uuid"
	// 	testingName := "My name"
	// 	returningError := fmt.Errorf("my custom error")

	// 	customer := entity.Customer{ID: testingUUID, Name: testingName}

	// 	customerMock.On("Find", testingUUID).Return(&customer, returningError)

	// 	uccfh := UseCaseCustomerFindHandler{customerRepository: &customerMock}

	// 	icfd := InputCustomerFindDTO{ID: testingUUID}

	// 	result, err := uccfh.Handle(icfd)

	// 	assert.Equal(t, err, returningError)
	// 	assert.Equal(t, result.ID, "")
	// 	assert.Equal(t, result.Name, "")
	// })
}
