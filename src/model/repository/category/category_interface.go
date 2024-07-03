package repository

import (
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"go.mongodb.org/mongo-driver/mongo"
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
}
