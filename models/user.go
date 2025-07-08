package models

type User struct {
	BaseModel
	Username string  `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string  `json:"password" gorm:"type:varchar(255);not null"`
	Email    string  `json:"email" gorm:"type:varchar(255);unique;not null"`
	Role     string  `json:"role" gorm:"type:varchar(50);default:'user'"`
	Carts    []Cart  `json:"carts,omitempty"`
	Orders   []Order `json:"orders,omitempty"`
}
