package repository

import (
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func (cr *categoryRepositoryInterface) FindCategory(categoryID string) (*entity.CategoryEntity, *rest_err.RestErr) {

	categoryCollection := os.Getenv(CATEGORY_COLLECTION)

	collection := cr.databaseConn.Collection(categoryCollection)
	var category entity.CategoryEntity
	filter := bson.D{{Key: "id", Value: categoryID}}

	err := collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		logger.Error(err.Error(), err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return &category, nil
}
