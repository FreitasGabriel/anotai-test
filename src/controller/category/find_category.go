package controller

import (
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/queue"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) FindCategoryByID(c *gin.Context) {

	category_id := c.Query("id")

	result, err := cc.service.FindCategory(category_id)
	if err != nil {
		logger.Error("error to find category", err, zap.String("journey", "findCategory"))
		c.JSON(http.StatusNotFound, "error to find category")
		return
	}

	queueErr := queue.QueueSendMessage(fmt.Sprintf(`{"owner": "%s"}`, result.OwnerID))
	if queueErr != nil {
		logger.Error("error to publish message on queue", queueErr, zap.String("journey", "createCategory"))
	}

	c.JSON(http.StatusOK, result)
}
