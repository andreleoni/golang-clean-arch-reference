package list

import (
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerListHandler struct {
	customerRepository repository.CustomerRepository
}

func NewUseCaseCustomerListHandler(rc repository.CustomerRepository) UseCaseCustomerListHandler {
	return UseCaseCustomerListHandler{customerRepository: rc}
}

func (uccfh UseCaseCustomerListHandler) Handle(icfd InputCustomerListDTO) (OutputCustomerListDTO, error) {
	response := OutputCustomerListDTO{}

	result, err := uccfh.customerRepository.List()
	if err != nil {
		return response, err
	}

	for _, v := range result {
		response.Customers = append(response.Customers, OutputCustomerDTO{ID: v.ID, Name: v.Name})
	}

	return response, nil
}
