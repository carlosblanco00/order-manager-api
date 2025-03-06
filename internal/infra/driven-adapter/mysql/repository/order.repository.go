package repository

import (
	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func (r OrderRepository) Create(order *model.Order) (*model.Order, error) {

	result := r.Db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (r *OrderRepository) GetById(id int) (*model.Order, error) {
	var order model.Order
	result := r.Db.Preload("OrderItems").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
