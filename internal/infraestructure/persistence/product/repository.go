package customer

import (
	"golang-clean-arch-reference/internal/domain/product/entity"

	"gorm.io/gorm"
)

type Product struct {
	sqlite *gorm.DB
}

func NewProduct(sqlite *gorm.DB) Product {
	return Product{sqlite: sqlite}
}

func (c Product) List() ([]*entity.Product, error) {
	product := []*entity.Product{}
	result := c.sqlite.Find(&product)

	return product, result.Error
}

func (c Product) Find(ID string) (*entity.Product, error) {
	product := entity.Product{ID: ID}
	result := c.sqlite.First(&product)

	return &product, result.Error
}

func (c Product) Create(ep *entity.Product) error {
	return c.sqlite.Create(&ep).Error
}

func (c Product) Update(ep *entity.Product) error {
	return c.sqlite.Save(&ep).Error
}

func (c Product) Enable(ep *entity.Product) error {
	ep.Status = "enabled"

	return c.Update(ep)
}

func (c Product) Disable(ep *entity.Product) error {
	ep.Status = "disabled"

	return c.Update(ep)
}
