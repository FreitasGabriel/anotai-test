package repository

import (
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/mongo"
)

func (pr *productRepositoryInterface) CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, error) {

	colletion := pr.databaseConn.Collection(COLLECTION_NAME)

	value := converter.ProductDomainToEntity(product)

	result, err := colletion.InsertOne(ctx, &value)
	if err != nil {
		return nil, err
	}
	return result, nil
}
