package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
)

type ProductService interface {
	BaseService[models.Product]
}

type productService struct {
	BaseService[models.Product]
}

func NewProductService(repo repository.ProductRepository) ProductService {
	base := NewBaservice(repo)
	return &productService{BaseService: base}
}
