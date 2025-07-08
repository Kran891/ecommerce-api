package services

import (
	"ecommerce-api/models"
	"ecommerce-api/repository"
)

type CartService interface {
	BaseService[models.Cart]
}

type cartService struct {
	BaseService[models.Cart]
}

func NewCartService(repo repository.CartRepository) CartService {
	base := NewBaservice(repo)
	return &cartService{BaseService: base}

}
