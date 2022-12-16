package update

import (
	"golang-clean-arch-reference/internal/domain/order/entity"
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	orderpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseOrderUpdateHandler_Integration(t *testing.T) {
	postgres.PGSetup()

	orderRepository := orderpersistence.NewOrder(postgres.PG)

	expectedID := "oder-uuid-2"
	expectedProductID := "product-uuid-1"
	expectedCustomerID := "customer-uuid-1"
	expectedQuantity := uint64(2)

	order := entity.Order{
		ID:         expectedID,
		ProductID:  expectedProductID,
		CustomerID: expectedCustomerID,
		Quantity:   uint64(expectedQuantity),
	}

	err := orderRepository.Create(&order)
	assert.Nil(t, err)

	t.Run("when updating the record", func(t *testing.T) {
		customerUpdateHandler := NewUseCaseOrderUpdateHandler(orderRepository)

		myNewQuantity := uint64(4)

		icfd := InputOrderUpdateDTO{
			ID:         expectedID,
			ProductID:  expectedProductID,
			CustomerID: expectedCustomerID,
			Quantity:   uint64(myNewQuantity),
		}

		result, err := customerUpdateHandler.Handle(icfd)
		assert.Nil(t, err)

		findResult, err := orderRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, findResult.ID, "")
		assert.Equal(t, findResult.ProductID, expectedProductID)
		assert.Equal(t, findResult.CustomerID, expectedCustomerID)
		assert.Equal(t, findResult.Quantity, uint64(4))
	})
}
