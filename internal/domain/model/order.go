package model

import "time"

type Order struct {
	ID           *uint      `json:"id" gorm:"primaryKey"`
	CustomerName string     `json:"customer_name" gorm:"column:customer_name;not null"`
	TotalAmount  float64    `json:"total_amount" gorm:"column:total_amount;not null"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:onDelete:CASCADE" json:"order_items"`
}

type OrderGateway interface {
	Create(order *Order) (*Order, error)
	GetById(id int) (*Order, error)
}
