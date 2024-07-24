package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	repository "github.com/FreitasGabriel/anotai-test/src/model/repository/category"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	service "github.com/FreitasGabriel/anotai-test/src/model/service/product"
)

func NewCategoryService(
	repository repository.CategoryRepositoryInterface,
	productService service.ProductDomainService) CategoryServiceInterface {
	return &categoryServiceInterface{
		repository:     repository,
		productService: productService,
	}
}

type categoryServiceInterface struct {
	repository     repository.CategoryRepositoryInterface
	productService service.ProductDomainService
}

type CategoryServiceInterface interface {
	CreateCategory(category domain.CategoryDomainInterface) *rest_err.RestErr
	DeleteCategory(id string) *rest_err.RestErr
	FindCategory(id string) (*entity.CategoryEntity, *rest_err.RestErr)
}
