package find

import (
	"fmt"
	"golang-clean-arch-reference/internal/domain/product/repository"
)

type UseCaseProductFindHandler struct {
	productRepository repository.ProductRepository
}

func NewUseCaseProductFindHandler(rc repository.ProductRepository) UseCaseProductFindHandler {
	return UseCaseProductFindHandler{productRepository: rc}
}

func (uccfh UseCaseProductFindHandler) Handle(ipfd InputProductFindDTO) (OutputProductFindDTO, error) {
	response := OutputProductFindDTO{}

	result, err := uccfh.productRepository.Find(ipfd.ID)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "record not found" {
			return response, nil
		}

		return response, err
	}

	response.ID = result.ID
	response.Name = result.Name
	response.Status = result.Status
	response.Price = result.Price

	return response, nil
}
