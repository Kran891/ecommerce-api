package routes

import (
	"ecommerce-api/controllers"
	"ecommerce-api/middleware"
	"ecommerce-api/repository"
	"ecommerce-api/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewAuthController(userService)
	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		//	auth.POST("/logout", controllers.Logout)
		auth.GET("/find/id/:id", userController.Find)
		auth.DELETE("/delete/id/:id", userController.Delete)
		auth.POST("/update", userController.Update)
		auth.POST("/logout", userController.Logout)
		auth.GET("/cartitems/:id", middleware.AuthMiddleware(), userController.CartItems)
	}

}
