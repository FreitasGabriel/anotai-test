package repository

import (
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (cr *categoryRepositoryInterface) FindCategory(categoryID string) (*entity.CategoryEntity, error) {

	category_colelction := os.Getenv(MONGO_COLLECTION)

	collection := cr.databaseConn.Collection(category_colelction)
	var category entity.CategoryEntity
	filter := bson.D{{Key: "id", Value: categoryID}}

	err := collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		logger.Error("error to find the document", err)
		return nil, err
	}

	return &category, nil
}
