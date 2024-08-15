package controller

import (
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/queue"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) DeleteCategory(c *gin.Context) {
	category_id := c.Param("id")

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

	queueErr := queue.QueueSendMessage(`{"owner": "1"}`)
	if queueErr != nil {
		logger.Error("error to publish message on queue", queueErr, zap.String("journey", "createCategory"))
	}

	c.JSON(http.StatusOK, "category deleted successfully")
}
