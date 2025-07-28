package models

import "github.com/google/uuid"

type Order struct {
	BaseModel
	UserID     uuid.UUID   `json:"userID"`
	OrderItems []OrderItem `json:"orderItems" gorm:"foreignKey:OrderID"`
	TotalPrice float64     `json:"totalPrice"`
}
