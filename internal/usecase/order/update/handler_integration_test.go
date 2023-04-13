package update

import (
	"golang-clean-arch-reference/internal/domain/order/entity"
	"golang-clean-arch-reference/internal/domain/valueobject"
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	orderpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseOrderUpdateHandler_Integration(t *testing.T) {
	sqlite.SQLiteSetup()

	orderRepository := orderpersistence.NewOrder(sqlite.Sqlite)

	expectedID := "oder-uuid-2"
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

	t.Run("when updating the record", func(t *testing.T) {
		customerUpdateHandler := NewUseCaseOrderUpdateHandler(orderRepository)

		myNewQuantity := uint64(4)

		icfd := InputOrderUpdateDTO{
			ID:         expectedID,
			ProductID:  expectedProductID,
			CustomerID: expectedCustomerID,
			Quantity:   uint64(myNewQuantity),
			Address: valueobject.Address{
				Street:     "anoter street",
				Number:     "333",
				Complement: "another complement",
				Zipcode:    "89219000",
				City:       "AnotherCity",
				Province:   "AnotherProvince",
			},
		}

		result, err := customerUpdateHandler.Handle(icfd)
		assert.Nil(t, err)

		findResult, err := orderRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, findResult.ID, "")
		assert.Equal(t, findResult.ProductID, expectedProductID)
		assert.Equal(t, findResult.CustomerID, expectedCustomerID)
		assert.Equal(t, findResult.Quantity, uint64(4))
		assert.Equal(t, findResult.Address.Street, icfd.Address.Street)
		assert.Equal(t, findResult.Address.Number, icfd.Address.Number)
		assert.Equal(t, findResult.Address.Complement, icfd.Address.Complement)
		assert.Equal(t, findResult.Address.Zipcode, icfd.Address.Zipcode)
		assert.Equal(t, findResult.Address.City, icfd.Address.City)
		assert.Equal(t, findResult.Address.Province, icfd.Address.Province)
	})
}
