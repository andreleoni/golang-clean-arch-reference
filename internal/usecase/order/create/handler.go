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

func (uccch UseCaseOrderCreateHandler) Handle(iocd InputOrderCreateDTO) (OutputOrderCreateDTO, error) {
	response := OutputOrderCreateDTO{}

	orderFactory := factory.NewOrderFactory()
	order := orderFactory.Create(iocd.ProductID, iocd.CustomerID, iocd.Quantity)

	order.Validate()

	if order.HasErrors() {
		return response, order.ErrorMessage()
	}

	err := uccch.orderRepository.Create(order)
	if err != nil {
		return response, err
	}

	response.ID = order.ID
	response.ProductID = order.ProductID
	response.CustomerID = order.CustomerID
	response.Quantity = order.Quantity

	return response, nil
}
