package models

import "github.com/google/uuid"

type Order struct {
	BaseModel
	UserID     uuid.UUID   `json:"userID"`
	OrderItems []OrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
}
