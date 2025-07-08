package repository

import "ecommerce-api/models"

type CartRepository interface {
	BaseRepository[models.Cart]
}

type cartRepository struct {
	GormRepository[models.Cart]
}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}
