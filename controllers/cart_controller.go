package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
)

type CartController interface {
	BaseController[models.Cart]
}
type cartController struct {
	BaseController[models.Cart]
}

func NewCartController(service services.CartService) CartController {
	prelods := []string{}
	base := NewBaseController(service, prelods)
	return &cartController{BaseController: base}
}
