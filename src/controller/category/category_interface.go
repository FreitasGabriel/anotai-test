package controller

import (
	service "github.com/FreitasGabriel/anotai-test/src/model/service/category"
	"github.com/gin-gonic/gin"
)

func NewCategoryController(service service.CategoryServiceInterface) CategoryControllerInterface {
	return &categoryControllerInterface{
		service: service,
	}
}

type categoryControllerInterface struct {
	service service.CategoryServiceInterface
}

type CategoryControllerInterface interface {
	CreateCategory(c *gin.Context)
	FindCategory(c *gin.Context)
}
