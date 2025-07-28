package models

import "github.com/google/uuid"

type OrderItem struct {
	BaseModel
	OrderID   uuid.UUID `json:"orderId" gorm:"index:idx_order_product,unique;not null"`
	ProductID uuid.UUID `json:"productId" gorm:"index:idx_order_product,unique;not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Product   Product
	Price     float64
	Order     Order
}
