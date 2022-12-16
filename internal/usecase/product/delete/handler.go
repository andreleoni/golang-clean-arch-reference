package delete

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerDeleteHandler struct {
	customerRepository repository.CustomerRepository
}

func NewUseCaseCustomerDeleteHandler(rc repository.CustomerRepository) UseCaseCustomerDeleteHandler {
	return UseCaseCustomerDeleteHandler{customerRepository: rc}
}

func (uccuh UseCaseCustomerDeleteHandler) Handle(icfd InputCustomerDeleteDTO) (OutputCustomerDeleteDTO, error) {
	response := OutputCustomerDeleteDTO{}

	customer, err := uccuh.customerRepository.Find(icfd.ID)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return response, fmt.Errorf("customer not found")
		}

		return response, err
	}

	err = uccuh.customerRepository.Delete(customer)
	if err != nil {
		return response, err
	}

	return response, nil
}
