package repository

import (
	"ecommerce-api/config"
	"ecommerce-api/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	BaseRepository[models.User]
	FindByEmail(email string) (user *models.User, err error)
	CartItems(id uuid.UUID) (user *models.User, err error)
	OrderItems(id uuid.UUID) (user *models.User, err error)
}

type userRepository struct {
	GormRepository[models.User]
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
func (r *userRepository) CartItems(id uuid.UUID) (user *models.User, err error) {
	err = config.DB.Preload("Carts").Preload("Carts.Product").First(&user, "id = ?", id).Error
	return

}
func (r *userRepository) OrderItems(id uuid.UUID) (user *models.User, err error) {
	err = config.DB.
		Preload("Orders").
		Preload("Orders.OrderItems").
		Preload("Orders.OrderItems.Product").
		First(&user, "id = ?", id).Error
	return
}
func NewUserRepository() UserRepository {
	return &userRepository{}
}
