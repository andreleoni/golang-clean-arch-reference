package create

import (
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	productpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseProductCreateHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	productRepository := productpersistence.NewProduct(sqlite.Sqlite)

	expectedName := "my name"
	expectedStatus := "enabled"
	expectedPrice := uint64(123)

	t.Run("when creating the handler", func(t *testing.T) {
		ipcd := InputProductCreateDTO{
			Name:   expectedName,
			Status: expectedStatus,
			Price:  expectedPrice,
		}

		producCreateHandler := NewUseCaseProductCreateHandler(productRepository)

		expectedResult := OutputProductCreateDTO{Name: expectedName, Status: expectedStatus, Price: expectedPrice}

		result, err := producCreateHandler.Handle(ipcd)

		assert.Nil(t, err)
		assert.NotEqual(t, result.ID, "")
		assert.Equal(t, expectedResult.Name, result.Name)
		assert.Equal(t, expectedResult.Status, result.Status)
		assert.Equal(t, expectedResult.Price, result.Price)

		findResult, err := productRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.Equal(t, result.ID, findResult.ID)
		assert.Equal(t, expectedResult.Name, result.Name)
		assert.Equal(t, expectedResult.Status, result.Status)
		assert.Equal(t, expectedResult.Price, result.Price)
	})
}
