package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.uber.org/zap"
)

func (ps *productDomainService) FindProductByID(id string) (*entity.ProductEntity, error) {

	product, err := ps.productRepository.FindProductByID(id)
	if err != nil {
		logger.Error("errot to find product", err, zap.String("journey", "findProductByID"))
		return nil, err
	}
	logger.Info("product find successfully")
	return product, nil
}

func (ps *productDomainService) FindProductsByTitle(title string) (*[]entity.ProductEntity, error) {
	products, err := ps.productRepository.FindProductsByTitle(title)
	if err != nil {
		logger.Error("error to find product with this title", err, zap.String("journey", "findProductsByID"))
		return nil, err
	}

	logger.Info("products find successfully")
	return products, nil
}
