package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/repository"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	repo := repository.NewCartRepository()
	service := services.NewCartService(repo)
	controller := controllers.NewCartController(service)
	cart := r.Group("/cart/items")
	{
		cart.POST("/", controller.Create)
		cart.DELETE("/:id", controller.Delete)
		cart.POST("/update", controller.Update)
	}
}
