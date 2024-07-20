package repository

import (
	"fmt"
	"os"

	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity/converter"
)

func (cr *categoryRepositoryInterface) CreateCategory(category domain.CategoryDomainInterface) error {

	mongo_collection := os.Getenv(CATEGORY_COLLECTION)
	collection := cr.databaseConn.Collection(mongo_collection)

	fmt.Println("repo", category)
	value := converter.CategoryDomainToEntity(category)

	fmt.Println("repo value", value)

	result, err := collection.InsertOne(ctx, value)
	if err != nil {
		return err
	}

	fmt.Println("result", result)

	return nil
}
