package service

import (
	"fmt"

	categoryDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
)

func (cr *categoryServiceInterface) CreateCategory(category categoryDomain.CategoryDomainInterface) error {

	fmt.Println("service", category)
	if err := cr.repository.CreateCategory(category); err != nil {
		return err
	}

	return nil
}
