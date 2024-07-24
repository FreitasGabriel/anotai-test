package service

import (
	"fmt"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.uber.org/zap"
)

func (ps *productDomainService) FindProductByID(id string) (*entity.ProductEntity, *rest_err.RestErr) {

	product, err := ps.productRepository.FindProductByID(id)
	if err != nil {
		logger.Error("errot to find product", err, zap.String("journey", "findProductByID"))
		return nil, &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}
	logger.Info("product find successfully")
	return product, nil
}

func (ps *productDomainService) FindProductsByTitle(title string) (*[]entity.ProductEntity, *rest_err.RestErr) {
	products, err := ps.productRepository.FindProductsByTitle(title)
	if err != nil {
		logger.Error("error to find product with this title", err, zap.String("journey", "findProductsByID"))
		return nil, &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	logger.Info("products find successfully")
	return products, nil
}

func (ps *productDomainService) FindProductsByCategoryID(category_id string) (*[]entity.ProductEntity, *rest_err.RestErr) {
	products, err := ps.productRepository.FindProductsByCategoryID(category_id)
	if err != nil {
		logger.Error(fmt.Sprintf("could not possible to find products with this category id : %s", category_id), err, zap.String("journey", "findProductsByProductID"))
		return nil, &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	logger.Info("products with this category_id finded successfully")
	return products, nil
}
