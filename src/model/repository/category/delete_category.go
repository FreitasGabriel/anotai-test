package repository

import (
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (cr *categoryRepositoryInterface) DeleteCategory(id string) error {

	collectionName := os.Getenv(CATEGORY_COLLECTION)
	collection := cr.databaseConn.Collection(collectionName)
	filter := bson.D{{Key: "id", Value: id}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error("could not to delete the category from database", err, zap.String("journey", "deleteCategory"))
		return err
	}

	logger.Info("category deleted succesfully")
	return nil
}
