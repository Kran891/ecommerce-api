package models

import "github.com/gofrs/uuid"

type Cart struct {
	BaseModel
	UserID    uuid.UUID `json:"userId" gorm:"not null;index:idx_user_product,unique"`
	ProductID uuid.UUID `json:"productId" gorm:"not null;index:id_user_product,unique"`
	Quantity  int64     `json:"quantity"`
	Product   Product
}
