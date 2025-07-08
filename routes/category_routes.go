package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/repository"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

func CaegoryRoutes(r *gin.Engine) {
	repo := repository.NewCategoryRepository()
	service := services.NewCategoryService(repo)
	controller := controllers.NewCategoryController(service)
	category := r.Group("/category")
	{
		category.POST("/create", controller.Create)
		category.POST("/update", controller.Update)
		category.GET("/categories", controller.FindAll)
		category.GET("/id/:id", controller.FindByID)
		category.DELETE("/id/:id", controller.Delete)
	}
}
