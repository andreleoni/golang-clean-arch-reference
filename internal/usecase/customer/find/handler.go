package find

import (
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerFindHandler struct {
	customerRepository repository.CustomerRepository
}

func NewUseCaseCustomerFindHandler(rc repository.CustomerRepository) UseCaseCustomerFindHandler {
	return UseCaseCustomerFindHandler{customerRepository: rc}
}

func (uccfh UseCaseCustomerFindHandler) Handle(icfd InputCustomerFindDTO) (OutputCustomerFindDTO, error) {
	response := OutputCustomerFindDTO{}

	result, err := uccfh.customerRepository.Find(icfd.ID)
	if err != nil {
		if err.Error() == "record not found" {
			return response, nil
		}

		return response, err
	}

	response.ID = result.ID
	response.Name = result.Name

	return response, nil
}
