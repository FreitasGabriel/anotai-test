package service

import (
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
	CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, error)
	FindProductByID(id string) (*entity.ProductEntity, error)
}
