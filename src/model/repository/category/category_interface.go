package repository

import (
	"context"

	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	CATEGORY_COLLECTION = "CATEGORY_COLLECTION_NAME"
	ctx                 = context.Background()
)

func NewCategoryRepository(database *mongo.Database) CategoryRepositoryInterface {
	return &categoryRepositoryInterface{
		databaseConn: database,
	}
}

type categoryRepositoryInterface struct {
	databaseConn *mongo.Database
}

type CategoryRepositoryInterface interface {
	CreateCategory(category domain.CategoryDomainInterface) error
	FindCategory(id string) (*entity.CategoryEntity, error)
	DeleteCategory(id string) *rest_err.RestErr
}
