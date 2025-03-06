package model

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Price     float64   `gorm:"column:price;not null" json:"price"`
	Stock     int       `gorm:"column:stock;not null" json:"stock"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type ProductGateway interface {
	Create(product *Product) (*Product, error)
	GetById(id int) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
}
