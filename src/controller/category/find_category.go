package controller

import (
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) FindCategoryByID(c *gin.Context) {

	category_id := c.Query("id")

	result, err := cc.service.FindCategory(category_id)
	if err != nil {
		logger.Error("error to find category", err, zap.String("journey", "findCategory"))
		c.JSON(http.StatusNotFound, "error to find category")
	}

	c.JSON(http.StatusOK, result)
}
