package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
)

type CategoryService interface {
	BaseService[models.Category]
}

type categoryService struct {
	BaseService[models.Category]
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	base := NewBaservice(repo)
	return &categoryService{BaseService: base}

}
