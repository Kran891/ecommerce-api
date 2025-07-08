package controllers

import (
	"ecommerce-api/logger"
	"ecommerce-api/services"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BaseController[T any] interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	BIND_ERROR     = "Unable to Bind %v to the model"
	DB_ERROR       = "Unable to Insert %v to the DB"
	UUID_ERROR     = "Invalid UUID Provided"
	ITEM_NOT_FOUND = "No %v found by the id %v"
	UPDATE_ERROR   = "Unable to Insert %v to the DB"
	DELETE_ERROR   = "Unable to Delete the %v with the id: %d"
)

type BaseControllerImpl[T any] struct {
	service  services.BaseService[T]
	preloads []string
}

func (bc *BaseControllerImpl[T]) Create(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		logger.Error(fmt.Sprintf(BIND_ERROR, reflect.TypeOf(item).Name()), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Provided"})
		return
	}

	if err := bc.service.Create(&item); err != nil {
		logger.Error(fmt.Sprintf(DB_ERROR, reflect.TypeOf(item).Name()), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to Create the Object"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (bc *BaseControllerImpl[T]) FindByID(c *gin.Context) {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.Error(UUID_ERROR, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id Provided"})
		return
	}
	item, err := bc.service.FindByID(id, bc.preloads...)
	if err != nil {
		logger.Error("DB Error", err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf(ITEM_NOT_FOUND, reflect.TypeOf(item).Name(), id))
		return
	}
	c.JSON(http.StatusFound, item)

}

func (bc *BaseControllerImpl[T]) FindAll(c *gin.Context) {
	items, err := bc.service.FindAll()
	if err != nil {
		logger.Error("DB Error", err)
		c.JSON(http.StatusInternalServerError, "Unable to process your request at the moment")
		return
	}
	c.JSON(http.StatusFound, items)
}

func (bc *BaseControllerImpl[T]) Update(c *gin.Context) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		logger.Error(fmt.Sprintf(BIND_ERROR, reflect.TypeOf(item).Name()), err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Provided"})
		return
	}
	if err := bc.service.Update(&item); err != nil {
		logger.Error(fmt.Sprintf(UPDATE_ERROR, reflect.TypeOf(item).Name()), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to process your request at the moment"})
		return
	}
	c.JSON(http.StatusFound, item)

}

func (bc *BaseControllerImpl[T]) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.Error(UUID_ERROR, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id Provided"})
		return
	}
	if err := bc.service.Delete(id); err != nil {
		logger.Error(fmt.Sprintf(DELETE_ERROR, "WILL FILL", id), err)
		c.JSON(http.StatusInternalServerError, "Unable to process your request at the moment")
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Delete the id :%v", id))
}

func NewBaseController[T any](service services.BaseService[T], preloads []string) BaseController[T] {

	return &BaseControllerImpl[T]{service: service, preloads: preloads}
}
