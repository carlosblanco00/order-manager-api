package repository

import (
	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"gorm.io/gorm"
)

type ProductRepositpry struct {
	Db *gorm.DB
}

func (pr ProductRepositpry) Create(product *model.Product) (*model.Product, error) {

	result := pr.Db.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (pr ProductRepositpry) GetById(id int) (*model.Product, error) {

	var product model.Product
	result := pr.Db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (pr ProductRepositpry) Update(product *model.Product) error {

	result := pr.Db.Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr ProductRepositpry) GetAll() ([]*model.Product, error) {
	var products []*model.Product

	if err := pr.Db.Find(&products).Error; err != nil {
		return nil, model.ManageError(model.ErrNotFound)
	}

	return products, nil
}
