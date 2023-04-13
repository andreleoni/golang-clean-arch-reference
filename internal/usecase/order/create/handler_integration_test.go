package create

import (
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	orderpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseOrderCreateHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	ordersRepository := orderpersistence.NewOrder(sqlite.Sqlite)

	expectedProductID := "product-uuid-1"
	expectedCustomerID := "customer-uuid-1"
	expectedQuantity := uint64(2)

	t.Run("when creating the handler", func(t *testing.T) {
		icfd := InputOrderCreateDTO{
			ProductID:  expectedProductID,
			CustomerID: expectedCustomerID,
			Quantity:   uint64(expectedQuantity),
		}

		orderCreateHandler := NewUseCaseOrderCreateHandler(ordersRepository)

		result, err := orderCreateHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.NotEqual(t, result.ID, "")
		assert.Equal(t, result.ProductID, expectedProductID)
		assert.Equal(t, result.CustomerID, expectedCustomerID)
		assert.Equal(t, result.Quantity, uint64(2))

		findResult, err := ordersRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, findResult.ID, "")
		assert.Equal(t, findResult.ProductID, expectedProductID)
		assert.Equal(t, findResult.CustomerID, expectedCustomerID)
		assert.Equal(t, findResult.Quantity, uint64(2))
	})
}
