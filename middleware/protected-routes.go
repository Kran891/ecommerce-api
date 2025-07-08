package middleware

import "github.com/gin-gonic/gin"

func ProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{"user": user})
		})
	}
}
