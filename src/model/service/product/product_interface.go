package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	repository "github.com/FreitasGabriel/anotai-test/src/model/repository/product"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductService(productRepository repository.ProductRepositoryInterface) ProductDomainService {
	return &productDomainService{
		productRepository,
	}
}

type productDomainService struct {
	productRepository repository.ProductRepositoryInterface
}

type ProductDomainService interface {
	CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, *rest_err.RestErr)
	FindProductByID(id string) (*entity.ProductEntity, *rest_err.RestErr)
	FindProductsByTitle(title string) (*[]entity.ProductEntity, *rest_err.RestErr)
	FindProductsByCategoryID(category_id string) (*[]entity.ProductEntity, *rest_err.RestErr)
	DeleteProduct(id string) *rest_err.RestErr
}
