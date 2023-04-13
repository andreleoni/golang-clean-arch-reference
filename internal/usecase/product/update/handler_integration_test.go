package update

import (
	"golang-clean-arch-reference/internal/domain/product/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	productpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseProductUpdateHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	productRepository := productpersistence.NewProduct(sqlite.Sqlite)

	startProductID := "my-uuid"
	startProductName := "my name"
	startProductStatus := "enabled"
	StartProductPrice := uint64(123)

	product := entity.Product{ID: startProductID, Name: startProductName, Status: startProductStatus, Price: StartProductPrice}

	err := productRepository.Create(&product)
	assert.Nil(t, err)

	t.Run("when updating the record", func(t *testing.T) {
		productUpdateHandler := NewUseCaseProductUpdateHandler(productRepository)

		updateName := "my another name"
		updateStatus := "disabled"
		updatePrice := uint64(20)

		icfd := InputProductUpdateDTO{
			ID:     product.ID,
			Name:   updateName,
			Status: updateStatus,
			Price:  updatePrice,
		}

		result, err := productUpdateHandler.Handle(icfd)
		assert.Nil(t, err)

		findResult, err := productRepository.Find(result.ID)
		assert.Nil(t, err)

		assert.Equal(t, result.ID, findResult.ID)
		assert.Equal(t, result.Name, updateName)
		assert.Equal(t, result.Status, updateStatus)
		assert.Equal(t, result.Price, updatePrice)
	})
}
