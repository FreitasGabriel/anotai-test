package repository

import (
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(database *mongo.Database) ProductRepositoryInterface {
	return &productRepositoryInterface{
		databaseConn: database,
	}
}

type productRepositoryInterface struct {
	databaseConn *mongo.Database
}

type ProductRepositoryInterface interface {
	CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, error)
}
