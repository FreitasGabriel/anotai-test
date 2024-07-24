package repository

import (
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (cr *categoryRepositoryInterface) DeleteCategory(id string) *rest_err.RestErr {

	collectionName := os.Getenv(CATEGORY_COLLECTION)
	collection := cr.databaseConn.Collection(collectionName)
	filter := bson.D{{Key: "id", Value: id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error("could not to delete the category from database", err, zap.String("journey", "deleteCategory"))
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.DeletedCount == 0 {
		logger.Error("there is no document with this categoryId", rest_err.NewNotFoundError("there is no document with this categoryId"), zap.String("journey", "deleteCategory"))
		return rest_err.NewNotFoundError("there is no document with this categoryId")
	}

	logger.Info("category deleted succesfully")
	return nil
}
