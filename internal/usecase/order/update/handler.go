package update

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/order/repository"
)

type UseCaseOrderUpdateHandler struct {
	orderRepository repository.OrderRepository
}

func NewUseCaseOrderUpdateHandler(oc repository.OrderRepository) UseCaseOrderUpdateHandler {
	return UseCaseOrderUpdateHandler{orderRepository: oc}
}

func (uccuh UseCaseOrderUpdateHandler) Handle(icfd InputOrderUpdateDTO) (OutputOrderUpdateDTO, error) {
	response := OutputOrderUpdateDTO{}

	order, err := uccuh.orderRepository.Find(icfd.ID)
	if err != nil {
		return response, err
	}

	if order.ID == "" {
		return response, fmt.Errorf("order not found")
	}

	order.ProductID = icfd.ProductID
	order.CustomerID = icfd.CustomerID
	order.Quantity = icfd.Quantity
	order.Address = icfd.Address

	order.Validate()

	if order.HasErrors() {
		return response, order.ErrorMessage()
	}

	err = uccuh.orderRepository.Update(order)
	if err != nil {
		return response, err
	}

	response.ID = order.ID
	response.ProductID = order.ProductID
	response.CustomerID = order.CustomerID
	response.Quantity = order.Quantity
	response.Address = order.Address

	return response, nil
}
