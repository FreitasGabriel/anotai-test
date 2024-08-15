package controller

import (
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/queue"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (pc *productControlerInterface) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	uuidErr := uuid.Validate(productID)
	if uuidErr != nil {
		logger.Error("This is not a valid id", uuidErr, zap.String("journey", "deleteProduct"))
		c.JSON(http.StatusBadRequest, "this is not a valid id")
		return
	}

	err := pc.service.DeleteProduct(productID)
	if err != nil {
		logger.Error("was not possible to delete product", err, zap.String("journey", "deleteProduct"))
		c.JSON(http.StatusBadRequest, "was not possible to delete product")
		return
	}

	queueErr := queue.QueueSendMessage(`{"owner": "1"}`)
	if queueErr != nil {
		logger.Error("error to publish message on queue", queueErr, zap.String("journey", "createCategory"))
	}

	c.JSON(http.StatusOK, "product deleted succesfully")
}
