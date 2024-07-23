package controller

import (
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pc *productControlerInterface) FindProductByID(c *gin.Context) {

	productID := c.Param("id")

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
	productTitle := c.Param("title")

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

func (pc *productControlerInterface) FindProductsByCategoryID(c *gin.Context) {
	categoryID := c.Param("categoryID")

	uuidErr := uuid.Validate(categoryID)
	if uuidErr != nil {
		logger.Error("categoryID is a not valid id", uuidErr, zap.String("Journey", "findProductsByCategoryID"))
		c.JSON(http.StatusBadRequest, "categoryID is a invalided id")
		return
	}

	result, err := pc.service.FindProductsByCategoryID(categoryID)
	if err != nil {
		logger.Error(fmt.Sprintf("could not possible to find products with this category id : %s", categoryID), err, zap.String("journey", "findProductsByCategoryID"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
