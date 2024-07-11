package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ps *productDomainService) DeleteProduct(id string) (string, error) {
	_, err := ps.productRepository.FindProductByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("the product does not exist in database", err, zap.String("journey", "deleteProduct"))
			return "", err
		}
		logger.Error("error to with dabatase operations", err, zap.String("journey", "deleteProduct"))
		return "", err
	}

	result, deleteError := ps.productRepository.DeleteProduct(id)
	if deleteError != nil {
		logger.Error("error to delete product from database", deleteError, zap.String("journey", "deleteProduct"))
		return "", deleteError
	}

	return result, nil
}
