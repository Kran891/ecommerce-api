package repository

import "ecommerce-api/models"

type ProductRepository interface {
	BaseRepository[models.Product]
}

type productRepository struct {
	GormRepository[models.Product]
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
