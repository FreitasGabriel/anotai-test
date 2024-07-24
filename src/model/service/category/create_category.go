package service

import (
	"fmt"

	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	categoryDomain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
)

func (cr *categoryServiceInterface) CreateCategory(category categoryDomain.CategoryDomainInterface) *rest_err.RestErr {

	fmt.Println("service", category)
	if err := cr.repository.CreateCategory(category); err != nil {
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	return nil
}
