package repository

import (
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (pr *productRepositoryInterface) DeleteProduct(id string) *rest_err.RestErr {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := pr.databaseConn.Collection(collection_name)
	filter := bson.D{{Key: "id", Value: id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error("error to delete document", rest_err.NewInternalServerError(err.Error()), zap.String("journey", "deleteProduct"))
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.DeletedCount == 0 {
		logger.Error("there is no document with this productId", rest_err.NewNotFoundError("there is no document with this productId"), zap.String("journey", "deleteproduct"))
		return rest_err.NewNotFoundError("there is no document with this productId")
	}

	logger.Info("product deleted succesfully")

	return nil
}
