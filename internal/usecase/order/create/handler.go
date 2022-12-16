package create

import (
	"golang-clean-arch-reference/internal/domain/order/factory"
	"golang-clean-arch-reference/internal/domain/order/repository"
)

type UseCaseOrderCreateHandler struct {
	orderRepository repository.OrderRepository
}

func NewUseCaseOrderCreateHandler(or repository.OrderRepository) UseCaseOrderCreateHandler {
	return UseCaseOrderCreateHandler{orderRepository: or}
}

func (uccch UseCaseOrderCreateHandler) Handle(icfd InputOrderFindDTO) (OutputOrderFindDTO, error) {
	response := OutputOrderCreateDTO{}

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
