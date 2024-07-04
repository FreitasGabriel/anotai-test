package repository

import (
	"context"

	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MONGO_COLLECTION = "CATEGORY_COLLECTION_NAME"
	ctx              = context.Background()
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
}
