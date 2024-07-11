package routes

import (
	categoryController "github.com/FreitasGabriel/anotai-test/src/controller/category"
	productController "github.com/FreitasGabriel/anotai-test/src/controller/product"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.RouterGroup, productController productController.ProductControllerInterface) {
	r.POST("/product", productController.CreateProduct)
	r.DELETE("/product/id/:id", productController.DeleteProduct)
	r.GET("/product/id/:id", productController.FindProductByID)
	r.GET("/product/title/:title", productController.FindProductsByTitle)
	r.GET("/product/categoryID/:categoryID", productController.FindProductsByCategoryID)
}

func InitCategoryRoutes(r *gin.RouterGroup, categoryController categoryController.CategoryControllerInterface) {
	r.POST("/category", categoryController.CreateCategory)
	r.GET("/category", categoryController.FindCategoryByID)

}
