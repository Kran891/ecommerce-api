package controllers

import (
	"ecommerce-api/models"
	"ecommerce-api/services"
	"ecommerce-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	UserService services.UserService
}

func NewAuthController(us services.UserService) *AuthController {
	return &AuthController{UserService: us}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := ac.UserService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (ac *AuthController) CartItems(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Id Provide", "id": id})
		return
	}
	if user, err := ac.UserService.CartItems(uid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id Provide", "id": id})
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := ac.UserService.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Mail or Password"})
		return
	}
	setCookie(c, token)
	c.JSON(http.StatusFound, gin.H{"token": token})
}
func (ac *AuthController) Update(c *gin.Context) {
	user, flag := utils.BindJSON[models.User](c)
	if !flag {
		return
	}
	if err := ac.UserService.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "User Updated"})
}

func (ac *AuthController) Delete(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Id Provide", "id": id})
		return
	}
	if err := ac.UserService.Delete(uid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id Provide", "id": id})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "User Deleted"})
}
func (ac *AuthController) Find(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Id Provide", "id": id})
		return
	}
	if user, err := ac.UserService.Find(uid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id Provide", "id": id})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}

}

func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, "Log off successfull...")
}
func setCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 3600*24, "/", "", false, true)
}
