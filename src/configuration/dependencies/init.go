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

func InitProductDep(database *mongo.Database) productController.ProductControllerInterface {
	repo := productRepository.NewProductRepository(database)
	service := productService.NewProductService(repo)
	return productController.NewProductController(service)
}

func InitCategoryDep(database *mongo.Database) categoryController.CategoryControllerInterface {
	repo := categoryRepository.NewCategoryRepository(database)
	service := categoryService.NewCategoryService(repo)
	return categoryController.NewCategoryController(service)

}
