package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	productDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pd *productDomainService) CreateProduct(product productDomain.ProductDomainInterface) (*mongo.InsertOneResult, *rest_err.RestErr) {

	result, err := pd.productRepository.CreateProduct(product)
	if err != nil {
		logger.Error("error to create product", err, zap.String("journey", "createProduct"))
		return nil, &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	logger.Info("product created with succesfully")
	return result, nil
}
