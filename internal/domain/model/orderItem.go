package model

import "time"

type OrderItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int       `gorm:"column:order_id;not null" json:"order_id"`
	ProductID int       `gorm:"column:product_id;not null" json:"product_id"`
	Quantity  int       `gorm:"column:quantity;not null" json:"quantity"`
	Subtotal  float64   `gorm:"column:subtotal;not null" json:"subtotal"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Order   *Order   `gorm:"foreignKey:OrderID;constraint:onDelete:CASCADE" json:"order"`
	Product *Product `gorm:"foreignKey:ProductID;constraint:onDelete:CASCADE" json:"product"`
}
