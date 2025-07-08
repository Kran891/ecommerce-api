package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
)

type CategoryController interface {
	BaseController[models.Category]
}

type categoryController struct {
	BaseControllerImpl[models.Category]
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{
		BaseControllerImpl: BaseControllerImpl[models.Category]{service: service,
			preloads: []string{"Products"},
		},
	}
}
