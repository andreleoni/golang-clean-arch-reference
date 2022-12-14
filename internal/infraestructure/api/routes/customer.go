package routes

import (
	"golang-clean-arch-reference/internal/infraestructure/database/postgres"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	customerfind "golang-clean-arch-reference/internal/usecase/customer/find"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerID struct {
	ID string `uri:"id" binding:"required"`
}

func CustomerRoutes(routes *gin.Engine) {
	customerRepository := customer.NewCustomer(postgres.PG)

	routes.GET("/customer/:id", func(c *gin.Context) {
		customerID := CustomerID{}

		if err := c.ShouldBindUri(&customerID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		icfd := customerfind.InputCustomerFindDTO{ID: customerID.ID}

		customerFindUseCase := customerfind.NewUseCaseCustomerFindHandler(customerRepository)

		result, err := customerFindUseCase.Handle(icfd)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else if result.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": result.ID, "name": result.Name})
		}
	})
}
