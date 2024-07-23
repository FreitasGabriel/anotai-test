package configuration

import (
	categoryController "github.com/FreitasGabriel/anotai-test/src/controller/category"
	productController "github.com/FreitasGabriel/anotai-test/src/controller/product"
	categoryRepository "github.com/FreitasGabriel/anotai-test/src/model/repository/category"
	productRepository "github.com/FreitasGabriel/anotai-test/src/model/repository/product"
	categoryService "github.com/FreitasGabriel/anotai-test/src/model/service/category"
	productService "github.com/FreitasGabriel/anotai-test/src/model/service/product"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDep(database *mongo.Database) (
	productController.ProductControllerInterface,
	categoryController.CategoryControllerInterface) {

	productRepo := productRepository.NewProductRepository(database)
	productService := productService.NewProductService(productRepo)
	categoryRepo := categoryRepository.NewCategoryRepository(database)
	categoryService := categoryService.NewCategoryService(categoryRepo, productService)

	return productController.NewProductController(productService), categoryController.NewCategoryController(categoryService)
}
