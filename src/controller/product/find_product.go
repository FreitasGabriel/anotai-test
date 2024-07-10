package controller

import (
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pc *productControlerInterface) FindProductByID(c *gin.Context) {

	productID := c.Query("id")

	uuidErr := uuid.Validate(productID)
	if uuidErr != nil {
		logger.Error("productID is a not valid id", uuidErr, zap.String("kourney", "findProductByID"))
		c.JSON(http.StatusBadRequest, "ProductID is a invalided id")
	}

	result, err := pc.service.FindProductByID(productID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("product not found in database", err, zap.String("journey", "findProductByID"))
			c.JSON(http.StatusNotFound, err)
		}
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}

func (pc *productControlerInterface) FindProductsByTitle(c *gin.Context) {
	productTitle := c.Query("title")

	result, err := pc.service.FindProductsByTitle(productTitle)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("product not found in database", err, zap.String("journey", "findProductsByTitle"))
			c.JSON(http.StatusNotFound, err)
		}
		logger.Error("error to get products", err, zap.String("journey", "findProductsByTitle"))
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, result)
}
