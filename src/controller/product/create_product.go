package controller

import (
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/controller/model/request"
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (pc productControlerInterface) CreateProduct(c *gin.Context) {
	var product = &request.ProductRequest{}

	if err := c.ShouldBindJSON(&product); err != nil {
		logger.Error("error to bind json", err, zap.String("journey", "createProduct"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("error to marshal json"))
		return
	}

	productID := uuid.New().String()

	domain := domain.NewProductDomain(
		productID,
		product.Title,
		product.Description,
		product.CategoryID,
		product.OwnerID,
		product.Price,
	)

	_, err := pc.service.CreateProduct(domain)
	if err != nil {
		logger.Error("error to create product handler", err, zap.String("journey", "createProduct"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("error to create productr handler"))
		return
	}

	logger.Info("product created successfully", zap.String("journey", "createProduct"))
	c.JSON(http.StatusCreated, "product created successfully")

}
