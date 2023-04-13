package find

import (
	"golang-clean-arch-reference/internal/domain/order/entity"
	"golang-clean-arch-reference/internal/domain/valueobject"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	orderpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseOrderFindHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	orderRepository := orderpersistence.NewOrder(sqlite.Sqlite)

	expectedID := "oder-uuid"
	expectedProductID := "product-uuid-1"
	expectedCustomerID := "customer-uuid-1"
	expectedQuantity := uint64(2)

	order := entity.Order{
		ID:         expectedID,
		ProductID:  expectedProductID,
		CustomerID: expectedCustomerID,
		Quantity:   uint64(expectedQuantity),
		Address: valueobject.Address{
			Street:     "my street",
			Number:     "123",
			Complement: "my complement",
			Zipcode:    "89219333",
			City:       "Joinville",
			Province:   "SC",
		},
	}

	err := orderRepository.Create(&order)
	assert.Nil(t, err)

	t.Run("when finding the record", func(t *testing.T) {
		customerFindHandler := NewUseCaseOrderFindHandler(orderRepository)

		iofd := InputOrderFindDTO{ID: expectedID}

		result, err := customerFindHandler.Handle(iofd)
		assert.Nil(t, err)

		findResult, err := orderRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, findResult.ID, "")
		assert.Equal(t, findResult.ProductID, expectedProductID)
		assert.Equal(t, findResult.CustomerID, expectedCustomerID)
		assert.Equal(t, findResult.Quantity, uint64(2))
		assert.Equal(t, findResult.Address.Street, order.Address.Street)
		assert.Equal(t, findResult.Address.Number, order.Address.Number)
		assert.Equal(t, findResult.Address.Complement, order.Address.Complement)
		assert.Equal(t, findResult.Address.Zipcode, order.Address.Zipcode)
		assert.Equal(t, findResult.Address.City, order.Address.City)
		assert.Equal(t, findResult.Address.Province, order.Address.Province)
	})
}
