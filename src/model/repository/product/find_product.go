package repository

import (
	"fmt"
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pr *productRepositoryInterface) FindProductByID(id string) (*entity.ProductEntity, *rest_err.RestErr) {

	product := entity.ProductEntity{}
	collectionName := os.Getenv(COLLECTION_NAME)
	collection := pr.databaseConn.Collection(collectionName)
	filter := bson.D{{Key: "id", Value: id}}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("product not found", err, zap.String("journey", "findProductByID"))
			return nil, rest_err.NewNotFoundError("product not found")
		}
		logger.Error("error to get product", err, zap.String("journey", "findProductByID"))
		return nil, rest_err.NewInternalServerError("error to get product")
	}

	return &product, nil
}
func (pr *productRepositoryInterface) FindProductsByTitle(title string) (*[]entity.ProductEntity, *rest_err.RestErr) {

	products := []entity.ProductEntity{}
	collectionName := os.Getenv(COLLECTION_NAME)
	collection := pr.databaseConn.Collection(collectionName)
	filter := bson.D{{Key: "title", Value: bson.M{"$regex": fmt.Sprintf("^%s", title)}}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("product not found", err, zap.String("journey", "findProductByTitle"))
			return nil, rest_err.NewNotFoundError("product not found")
		}
		logger.Error("error to get product", err, zap.String("journey", "findProductByTitle"))
		return nil, rest_err.NewInternalServerError("error to get product")
	}

	defer cursor.Close(ctx)

	if cursorErr := cursor.All(ctx, &products); cursorErr != nil {
		logger.Error("could not possible to cursor into slice", err, zap.String("journey", "findProductByTitle"))
	}

	return &products, nil
}
func (pr *productRepositoryInterface) FindProductByCategoryID(category_id string) (*mongo.SingleResult, *rest_err.RestErr) {
	return nil, nil
}
