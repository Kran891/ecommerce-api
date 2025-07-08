package main

import (
	"ecommerce-api/config"
	"ecommerce-api/logger"
	"ecommerce-api/middleware"
	"ecommerce-api/routes"

	"github.com/gin-gonic/gin"
)

// @title           Ecommerce API
// @version         1.0
// @description     Monolithic Gin + GORM ecommerce backend.
// @contact.name    Kranthi
// @contact.email   kranthi.gavireddy.code@gmail.com
// @license.name    MIT
// @host            localhost:8080
// @BasePath        /
// @schemes         http
func main() {
	logger.Init()
	logger.Info("Starting the application...")
	r := gin.Default()
	r.Use(middleware.RequestLogger())
	config.ConncectDB()
	routes.Routes(r)
	middleware.ProtectedRoutes(r)
	r.Run(":8080")
}
