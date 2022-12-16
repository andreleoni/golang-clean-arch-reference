package find

import (
	"golang-clean-arch-reference/internal/domain/order/repository"
)

type UseCaseOrderFindHandler struct {
	orderRepository repository.OrderRepository
}

func NewUseCaseOrderFindHandler(or repository.OrderRepository) UseCaseOrderFindHandler {
	return UseCaseOrderFindHandler{orderRepository: or}
}

func (uccfh UseCaseOrderFindHandler) Handle(icfd InputOrderFindDTO) (OutputOrderFindDTO, error) {
	response := OutputOrderFindDTO{}

	order, err := uccfh.orderRepository.Find(icfd.ID)
	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return response, nil
		}

		return response, err
	}

	response.ID = order.ID
	response.ProductID = order.ProductID
	response.CustomerID = order.CustomerID
	response.Quantity = order.Quantity

	return response, nil
}
