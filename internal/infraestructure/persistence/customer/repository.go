package customer

import (
	"golang-clean-arch-reference/internal/domain/customer/entity"

	"gorm.io/gorm"
)

type Customer struct {
	sqlite *gorm.DB
}

func NewCustomer(sqlite *gorm.DB) Customer {
	return Customer{sqlite: sqlite}
}

func (c Customer) List() ([]*entity.Customer, error) {
	customer := []*entity.Customer{}
	result := c.sqlite.Find(&customer)

	return customer, result.Error
}

func (c Customer) Find(ID string) (*entity.Customer, error) {
	customer := entity.Customer{ID: ID}
	result := c.sqlite.First(&customer)

	return &customer, result.Error
}

func (c Customer) Create(ec *entity.Customer) error {
	result := c.sqlite.Create(&ec)

	return result.Error
}

func (c Customer) Update(ec *entity.Customer) error {
	result := c.sqlite.Save(&ec)

	return result.Error
}

func (c Customer) Delete(ec *entity.Customer) error {
	result := c.sqlite.Delete(&ec)

	return result.Error
}
