package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
)

type ProductController interface {
	BaseController[models.Product]
}

type productController struct {
	BaseControllerImpl[models.Product]
}

func NewProductController(service services.ProductService) ProductController {
	return &productController{BaseControllerImpl: BaseControllerImpl[models.Product]{service: service, preloads: []string{}}}
}
