package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/repository"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

func ProductRouts(r *gin.Engine) {
	repo := repository.NewProductRepository()
	service := services.NewProductService(repo)
	controller := controllers.NewProductController(service)
	product := r.Group("/product")
	{
		product.POST("/create", controller.Create)
		product.POST("/update", controller.Update)
		product.GET("/products", controller.FindAll)
		product.GET("/id/:id", controller.FindByID)
		product.DELETE("/id/:id", controller.Delete)
	}
}
