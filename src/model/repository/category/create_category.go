package repository

import (
	"fmt"
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity/converter"
)

func (cr *categoryRepositoryInterface) CreateCategory(category domain.CategoryDomainInterface) *rest_err.RestErr {

	mongo_collection := os.Getenv(CATEGORY_COLLECTION)
	collection := cr.databaseConn.Collection(mongo_collection)

	fmt.Println("repo", category)
	value := converter.CategoryDomainToEntity(category)

	fmt.Println("repo value", value)

	_, err := collection.InsertOne(ctx, value)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
