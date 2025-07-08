package repository

import "ecommerce-api/models"

type CategoryRepository interface {
	BaseRepository[models.Category]
}

type categoryRepository struct {
	GormRepository[models.Category]
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}
