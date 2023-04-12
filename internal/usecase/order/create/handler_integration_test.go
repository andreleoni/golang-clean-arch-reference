package create

import (
	"golang-clean-arch-reference/internal/domain/valueobject"
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	orderpersistence "golang-clean-arch-reference/internal/infraestructure/persistence/order"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseOrderCreateHandler_Integration(t *testing.T) {
	postgres.PGSetup()

	ordersRepository := orderpersistence.NewOrder(postgres.PG)

	expectedProductID := "product-uuid-1"
	expectedCustomerID := "customer-uuid-1"
	expectedQuantity := uint64(2)

	t.Run("when creating the handler", func(t *testing.T) {
		icfd := InputOrderCreateDTO{
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

		orderCreateHandler := NewUseCaseOrderCreateHandler(ordersRepository)

		result, err := orderCreateHandler.Handle(icfd)

		assert.Nil(t, err)
		assert.NotEqual(t, result.ID, "")

		findResult, err := ordersRepository.Find(result.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, findResult.ID, "")
		assert.Equal(t, findResult.ProductID, expectedProductID)
		assert.Equal(t, findResult.CustomerID, expectedCustomerID)
		assert.Equal(t, findResult.Quantity, uint64(2))
		assert.Equal(t, findResult.Address.Street, icfd.Address.Street)
		assert.Equal(t, findResult.Address.Number, icfd.Address.Number)
		assert.Equal(t, findResult.Address.Complement, icfd.Address.Complement)
		assert.Equal(t, findResult.Address.Zipcode, icfd.Address.Zipcode)
		assert.Equal(t, findResult.Address.City, icfd.Address.City)
		assert.Equal(t, findResult.Address.Province, icfd.Address.Province)
	})
}
