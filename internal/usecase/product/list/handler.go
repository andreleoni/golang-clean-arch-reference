package list

import (
	"golang-clean-arch-reference/internal/domain/product/repository"
)

type UseCaseProductListHandler struct {
	productRepository repository.ProductRepository
}

func NewUseCaseProductListHandler(rp repository.ProductRepository) UseCaseProductListHandler {
	return UseCaseProductListHandler{productRepository: rp}
}

func (uccfh UseCaseProductListHandler) Handle(ipfd InputProductListDTO) (OutputProductListDTO, error) {
	response := OutputProductListDTO{}

	result, err := uccfh.productRepository.List()
	if err != nil {
		return response, err
	}

	for _, v := range result {
		response.Products = append(response.Products, OutputProductDTO{ID: v.ID, Name: v.Name, Status: v.Status, Price: v.Price})
	}

	return response, nil
}
