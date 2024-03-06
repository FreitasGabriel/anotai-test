package service

import (
	"github.com/FreitasGabriel/anotai-test/src/model"
	"github.com/FreitasGabriel/anotai-test/src/repository/product"
	"github.com/google/uuid"
)

type ProductService struct {
	DB product.ProductDB
}

func CreateProductService(product *model.ProductPayload, db product.ProductDB) error {

	productID := uuid.New()
	product.ProductID = productID.String()

	errDb := db.Create(product)
	if errDb != nil {
		return errDb
	}
	return nil
}

func ListProductsService(db product.ProductDB) ([]model.ProductPayload, error) {
	result, err := db.List()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetProductService()
