package service

import (
	"github.com/FreitasGabriel/anotai-test/src/model/repository/entity"
	"github.com/google/uuid"
)

func (cs *categoryServiceInterface) FindCategory(id string) (*entity.CategoryEntity, error) {

	err := uuid.Validate(id)
	if err != nil {
		return nil, err
	}

	result, err := cs.repository.FindCategory(id)
	if err != nil {
		return nil, err
	}

	category := entity.CategoryEntity{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		OwnerID:     result.OwnerID,
	}

	return &category, nil
}
