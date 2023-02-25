package list

import (
	"golang-clean-arch-reference/internal/domain/product/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	productpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseProductListHandler_Integration(t *testing.T) {
	postgres.PGSetup()

	productRepository := productpersistence.NewProduct(postgres.PG)

	firstItemID := "my-uuid-1"
	firstItemName := "my name 1"
	firstItemStatus := "enabled"
	firstItemPrice := uint64(123)

	product1 := entity.Product{ID: firstItemID, Name: firstItemName, Status: firstItemStatus, Price: firstItemPrice}

	err := productRepository.Create(&product1)
	assert.Nil(t, err)

	secondItemID := "my-uuid-2"
	secondItemName := "my name 1"

	product2 := entity.Product{ID: secondItemID, Name: secondItemName}

	err = productRepository.Create(&product2)
	assert.Nil(t, err)

	t.Run("when have records on database", func(t *testing.T) {
		icfd := InputProductListDTO{}

		productListHandler := NewUseCaseProductListHandler(productRepository)

		result, err := productListHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, len(result.Products), 2)

		assert.Equal(t, result.Products[0].ID, firstItemID)
		assert.Equal(t, result.Products[0].Name, firstItemName)

		assert.Equal(t, result.Products[1].ID, secondItemID)
		assert.Equal(t, result.Products[1].Name, secondItemName)
	})
}
