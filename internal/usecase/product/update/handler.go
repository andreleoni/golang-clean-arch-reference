package update

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/product/repository"
)

type UseCaseProductUpdateHandler struct {
	productRepository repository.ProductRepository
}

func NewUseCaseProductUpdateHandler(rc repository.ProductRepository) UseCaseProductUpdateHandler {
	return UseCaseProductUpdateHandler{productRepository: rc}
}

func (uccuh UseCaseProductUpdateHandler) Handle(ipud InputProductUpdateDTO) (OutputProductUpdateDTO, error) {
	response := OutputProductUpdateDTO{}

	product, err := uccuh.productRepository.Find(ipud.ID)
	if err != nil {
		return response, err
	}

	if product.ID == "" {
		return response, fmt.Errorf("product not found")
	}

	product.Name = ipud.Name
	product.Status = ipud.Status
	product.Price = ipud.Price

	product.Validate()

	if product.HasErrors() {
		return response, product.ErrorMessage()
	}

	err = uccuh.productRepository.Update(product)
	if err != nil {
		return response, err
	}

	response.ID = product.ID
	response.Name = product.Name
	response.Status = product.Status
	response.Price = product.Price

	return response, nil
}
