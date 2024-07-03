package service

import (
	"fmt"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pd *productDomainService) CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, error) {

	result, err := pd.productRepository.CreateProduct(product)
	if err != nil {
		logger.Error("error to create product", err, zap.String("journey", "createProduct"))
		return nil, err
	}

	fmt.Println(result)

	return result, nil
}
