package create

import (
	"golang-clean-arch-reference/internal/domain/customer/factory"
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerCreateHandler struct {
	customerRepository repository.CustomerRepository
}

func NewUseCaseCustomerCreateHandler(rc repository.CustomerRepository) UseCaseCustomerCreateHandler {
	return UseCaseCustomerCreateHandler{customerRepository: rc}
}

func (uccch UseCaseCustomerCreateHandler) Handle(icfd InputCustomerCreateDTO) (OutputCustomerCreateDTO, error) {
	response := OutputCustomerCreateDTO{}

	customerFactory := factory.NewCustomerFactory()
	customer := customerFactory.Create("", icfd.Name)

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
