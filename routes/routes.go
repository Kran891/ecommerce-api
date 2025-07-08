package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	AuthRoutes(r)
	CaegoryRoutes(r)
	ProductRouts(r)
	CartRoutes(r)
}
