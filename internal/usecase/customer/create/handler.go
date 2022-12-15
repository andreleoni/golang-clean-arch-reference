package create

import (
	"golang-clean-arch-reference/internal/domain/customer/factory"
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerCreateHandler struct {
	customerRepository repository.Customer
}

func NewUseCaseCustomerCreateHandler(rc repository.Customer) UseCaseCustomerCreateHandler {
	return UseCaseCustomerCreateHandler{customerRepository: rc}
}

func (uccch UseCaseCustomerCreateHandler) Handle(icfd InputCustomerCreateDTO) (OutputCustomerCreateDTO, error) {
	response := OutputCustomerCreateDTO{}

	customer := factory.NewCustomer("", icfd.Name)

	customer.Validate()

	if customer.HasErrors() {
		return response, customer.ErrorMessage()
	}

	err := uccch.customerRepository.Create(customer)
	if err != nil {
		return response, err
	}

	response.ID = customer.ID
	response.Name = customer.Name

	return response, nil
}
