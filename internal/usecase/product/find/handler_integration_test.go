package find

import (
	"golang-clean-arch-reference/internal/domain/product/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	productpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseProductFindHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	productRepository := productpersistence.NewProduct(sqlite.Sqlite)

	expectedID := "my-uuid"
	expectedName := "my name"
	expectedStatus := "enabled"
	expectedPrice := uint64(123)

	product := entity.Product{ID: expectedID, Name: expectedName, Status: expectedStatus, Price: expectedPrice}

	err := productRepository.Create(&product)
	assert.Nil(t, err)

	t.Run("when have the ID on database", func(t *testing.T) {
		icfd := InputProductFindDTO{ID: expectedID}

		productFindHandler := NewUseCaseProductFindHandler(productRepository)

		expectedResult := OutputProductFindDTO{
			ID:     expectedID,
			Name:   expectedName,
			Status: expectedStatus,
			Price:  expectedPrice,
		}

		result, err := productFindHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("when have not the ID on database", func(t *testing.T) {
		icfd := InputProductFindDTO{ID: "not existing ID"}

		productFindHandler := NewUseCaseProductFindHandler(productRepository)

		expectedResult := OutputProductFindDTO{}

		result, err := productFindHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
