package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ps *productDomainService) DeleteProduct(id string) *rest_err.RestErr {
	_, err := ps.productRepository.FindProductByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("the product does not exist in database", err, zap.String("journey", "deleteProduct"))
			return &rest_err.RestErr{
				Code:    err.Code,
				Message: err.Message,
				Err:     err.Err,
			}
		}
		logger.Error("error to with dabatase operations", err, zap.String("journey", "deleteProduct"))
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	deleteError := ps.productRepository.DeleteProduct(id)
	if deleteError != nil {
		logger.Error("error to delete product from database", deleteError, zap.String("journey", "deleteProduct"))
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	logger.Info("product deleted succesfully")

	return nil
}
