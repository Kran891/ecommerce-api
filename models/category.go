package models

type Category struct {
	BaseModel
	Name     string    `json:"name" gorm:"unique;not null"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
