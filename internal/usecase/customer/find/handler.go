package find

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/customer/repository"
)

type UseCaseCustomerFindHandler struct {
	customerRepository repository.Customer
}

func NewUseCaseCustomerFindHandler(rc repository.Customer) UseCaseCustomerFindHandler {
	return UseCaseCustomerFindHandler{customerRepository: rc}
}

func (uccfh UseCaseCustomerFindHandler) Handle(icfd InputCustomerFindDTO) (OutputCustomerFindDTO, error) {
	response := OutputCustomerFindDTO{}

	result, err := uccfh.customerRepository.Find(icfd.ID)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "pg: no rows in result set" {
			return response, nil
		}

		return response, err
	}

	response.ID = result.ID
	response.Name = result.Name

	return response, nil
}
