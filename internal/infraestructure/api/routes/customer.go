package routes

import (
	"golang-clean-arch-reference/internal/infraestructure/database/sqlite"
	"golang-clean-arch-reference/internal/infraestructure/persistence/customer"
	customercreate "golang-clean-arch-reference/internal/usecase/customer/create"
	customerdelete "golang-clean-arch-reference/internal/usecase/customer/delete"
	customerfind "golang-clean-arch-reference/internal/usecase/customer/find"
	customerlist "golang-clean-arch-reference/internal/usecase/customer/list"
	customerupdate "golang-clean-arch-reference/internal/usecase/customer/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindCustomerID struct {
	ID string `uri:"id" binding:"required"`
}

type CreateCustomer struct {
	Name string `uri:"name" json:"name" binding:"required"`
}

type UpdateCustomer struct {
	Name string `uri:"name" json:"name" binding:"required"`
}

func CustomerRoutes(routes *gin.Engine) {
	customerRepository := customer.NewCustomer(sqlite.Sqlite)

	routes.GET("/customer/:id", func(c *gin.Context) {
		customerID := FindCustomerID{}

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

	routes.POST("/customer", func(c *gin.Context) {
		createCustomer := CreateCustomer{}

		if err := c.ShouldBind(&createCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		icfd := customercreate.InputCustomerCreateDTO{Name: createCustomer.Name}

		customerCreateUseCase := customercreate.NewUseCaseCustomerCreateHandler(customerRepository)

		result, err := customerCreateUseCase.Handle(icfd)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": result.ID, "name": result.Name})
		}
	})

	routes.GET("/customers", func(c *gin.Context) {
		icld := customerlist.InputCustomerListDTO{}

		customerListUseCase := customerlist.NewUseCaseCustomerListHandler(customerRepository)

		result, err := customerListUseCase.Handle(icld)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else if len(result.Customers) == 0 {
			c.JSON(http.StatusOK, gin.H{"customers": []string{}})
		} else {
			c.JSON(http.StatusOK, result)
		}
	})

	routes.PUT("/customer/:id", func(c *gin.Context) {
		customerID := FindCustomerID{}

		if err := c.ShouldBindUri(&customerID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		updateCustomer := UpdateCustomer{}

		if err := c.ShouldBind(&updateCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		icud := customerupdate.InputCustomerUpdateDTO{ID: customerID.ID, Name: updateCustomer.Name}

		cuuc := customerupdate.NewUseCaseCustomerUpdateHandler(customerRepository)

		result, err := cuuc.Handle(icud)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": result.ID, "name": result.Name})
		}
	})

	routes.DELETE("/customer/:id", func(c *gin.Context) {
		customerID := FindCustomerID{}

		if err := c.ShouldBindUri(&customerID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		icud := customerdelete.InputCustomerDeleteDTO{ID: customerID.ID}

		uccdh := customerdelete.NewUseCaseCustomerDeleteHandler(customerRepository)

		_, err := uccdh.Handle(icud)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"id": customerID.ID})
		}
	})
}
