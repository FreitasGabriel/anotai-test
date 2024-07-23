package controller

import (
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) DeleteCategory(c *gin.Context) {

	category_id := c.Param("id")

	fmt.Println("category", category_id)

	uuidError := uuid.Validate(category_id)
	if uuidError != nil {
		logger.Error("this is not a valid uuid", uuidError, zap.String("journey", "deleteCategory"))
		c.JSON(http.StatusBadRequest, "this is a not valid uuid")
		return
	}

	err := cc.service.DeleteCategory(category_id)
	if err != nil {
		logger.Error("error to delete category", err, zap.String("journey", "deleteCategory"))
		c.JSON(err.Code, "error to delete category")
		return
	}

	c.JSON(http.StatusOK, "category deleted successfully")

}
