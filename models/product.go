package models

import "github.com/gofrs/uuid"

type Product struct {
	BaseModel
	Name       string    `json:"name" gorm:"not null;index"`
	Quantity   int64     `json:"quantity" gorm:"column:quantity;type:bigint"`
	Price      float64   `json:"price"`
	CategoryID uuid.UUID `json:"categoryId" gorm:"column:category_id;type:uuid;not null"`
}
