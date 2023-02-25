package create

import (
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	productpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseProductCreateHandler_Integration(t *testing.T) {
	postgres.PGSetup()

	productRepository := productpersistence.NewProduct(postgres.PG)

	expectedName := "my name"

	t.Run("when creating the handler", func(t *testing.T) {
		ipcd := InputProductCreateDTO{Name: expectedName, Status: "enabled", Price: 123}

		customerFindHandler := NewUseCaseProductCreateHandler(productRepository)

		expectedResult := OutputProductCreateDTO{Name: expectedName}

		result, err := customerFindHandler.Handle(ipcd)

		assert.Nil(t, err)
		assert.NotEqual(t, result.ID, "")
		assert.Equal(t, expectedResult.Name, result.Name)

		findResult, err := productRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.Equal(t, result.ID, findResult.ID)
		assert.Equal(t, expectedResult.Name, result.Name)
	})
}
