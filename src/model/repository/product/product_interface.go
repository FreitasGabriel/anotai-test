package repository

import (
	"context"

	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

const COLLECTION_NAME = "PRODUCT_COLLECTION_NAME"

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
	FindProductByID(id string) (*entity.ProductEntity, *rest_err.RestErr)
	FindProductByTitle(title string) (*mongo.SingleResult, *rest_err.RestErr)
	FindProductByCategoryID(category_id string) (*mongo.SingleResult, *rest_err.RestErr)
}
