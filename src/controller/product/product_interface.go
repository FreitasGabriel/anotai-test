package controller

import (
	service "github.com/FreitasGabriel/anotai-test/src/model/service/product"
	"github.com/gin-gonic/gin"
)

func NewProductController(service service.ProductDomainService) ProductControllerInterface {
	return &productControlerInterface{
		service,
	}
}

type productControlerInterface struct {
	service service.ProductDomainService
}

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	FindProductByID(c *gin.Context)
	FindProductsByTitle(c *gin.Context)
}
