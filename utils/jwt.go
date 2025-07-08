package utils

import (
	"ecommerce-api/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MakeJWT(user *models.User) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"sub":      fmt.Sprint(user.BaseModel.ID), // subject = user‑id
		"username": user.Username,
		"role":     user.Role,
		"email":    user.Email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // << int64, not time.Time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func BindJSON[T any](c *gin.Context) (*T, bool) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return nil, false
	}
	return &input, true
}
