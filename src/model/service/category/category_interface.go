package service

import (
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	repository "github.com/FreitasGabriel/anotai-test/src/model/repository/category"
)

func NewCategoryService(repository repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &categoryServiceInterface{
		repository: repository,
	}
}

type categoryServiceInterface struct {
	repository repository.CategoryRepositoryInterface
}

type CategoryServiceInterface interface {
	CreateCategory(category domain.CategoryDomainInterface) error
}
