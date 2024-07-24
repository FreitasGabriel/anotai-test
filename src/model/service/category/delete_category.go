package service

import (
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (cs *categoryServiceInterface) DeleteCategory(id string) *rest_err.RestErr {

	result, err := cs.productService.FindProductsByCategoryID(id)
	if err != nil {
		logger.Error("error to find products with this categoryID", err, zap.String("jounery", "deleteCategory"))
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}
	if len(*result) > 0 {
		logger.Error(
			"could not possible to delete category because there are products with this categoryID",
			rest_err.NewForbiddenError("could not possible to delete category because there are products with this categoryID"),
			zap.String("journey", "deleteCategory"),
		)
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	deleteErr := cs.repository.DeleteCategory(id)
	if deleteErr != nil {
		logger.Error("Error to delete category from database", err, zap.String("journey", "deleteCategory"))
		return &rest_err.RestErr{
			Code:    err.Code,
			Message: err.Message,
			Err:     err.Err,
		}
	}

	logger.Info("Cataegory deleted succesfully")
	return nil
}
