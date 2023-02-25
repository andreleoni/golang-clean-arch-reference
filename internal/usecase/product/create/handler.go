package create

import (
	"golang-clean-arch-reference/internal/domain/product/factory"
	"golang-clean-arch-reference/internal/domain/product/repository"
)

type UseCaseProductCreateHandler struct {
	productRepository repository.ProductRepository
}

func NewUseCaseProductCreateHandler(rc repository.ProductRepository) UseCaseProductCreateHandler {
	return UseCaseProductCreateHandler{productRepository: rc}
}

func (uccch UseCaseProductCreateHandler) Handle(ipfd InputProductCreateDTO) (OutputProductCreateDTO, error) {
	response := OutputProductCreateDTO{}

	productFactory := factory.NewProductFactory()
	product := productFactory.Create(ipfd.Name, ipfd.Status, ipfd.Price)

	product.Validate()

	if product.HasErrors() {
		return response, product.ErrorMessage()
	}

	err := uccch.productRepository.Create(product)
	if err != nil {
		return response, err
	}

	response.ID = product.ID
	response.Name = product.Name
	response.Price = product.Price

	return response, nil
}
