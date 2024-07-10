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

	return product, nil
}
