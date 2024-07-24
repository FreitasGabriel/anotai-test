package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"github.com/google/uuid"
)

func (cs *categoryServiceInterface) FindCategory(id string) (*entity.CategoryEntity, *rest_err.RestErr) {

	err := uuid.Validate(id)
	if err != nil {
		return nil, &rest_err.RestErr{
			Code:    400,
			Message: err.Error(),
		}
	}

	result, repoErr := cs.repository.FindCategory(id)
	if repoErr != nil {
		return nil, &rest_err.RestErr{
			Code:    500,
			Message: err.Error(),
		}
	}

	category := entity.CategoryEntity{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		OwnerID:     result.OwnerID,
	}

	return &category, nil
}
