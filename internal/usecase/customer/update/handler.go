package update

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerUpdateHandler struct {
	customerRepository repository.Customer
}

func NewUseCaseCustomerUpdateHandler(rc repository.Customer) UseCaseCustomerUpdateHandler {
	return UseCaseCustomerUpdateHandler{customerRepository: rc}
}

func (uccuh UseCaseCustomerUpdateHandler) Handle(icfd InputCustomerUpdateDTO) (OutputCustomerUpdateDTO, error) {
	response := OutputCustomerUpdateDTO{}

	customer, err := uccuh.customerRepository.Find(icfd.ID)
	if err != nil {
		return response, err
	}

	if customer.ID == "" {
		return response, fmt.Errorf("customer not found")
	}

	customer.Name = icfd.Name

	customer.Validate()

	if customer.HasErrors() {
		return response, customer.ErrorMessage()
	}

	err = uccuh.customerRepository.Update(customer)
	if err != nil {
		return response, err
	}

	response.ID = customer.ID
	response.Name = customer.Name

	return response, nil
}
