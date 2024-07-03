package routes

import (
	categoryController "github.com/FreitasGabriel/anotai-test/src/controller/category"
	productController "github.com/FreitasGabriel/anotai-test/src/controller/product"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.RouterGroup, productController productController.ProductControllerInterface) {
	r.POST("/product", productController.CreateProduct)
}

func InitCategoryRoutes(r *gin.RouterGroup, categoryController categoryController.CategoryControllerInterface) {
	r.POST("/category", categoryController.CreateCategory)
}