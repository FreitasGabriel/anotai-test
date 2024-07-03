package controller

import (
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/controller/model/request"
	domain "github.com/FreitasGabriel/anotai-test/src/model/domain/category"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) CreateCategory(c *gin.Context) {
	var category request.CategoryRequest

	if err := c.ShouldBindJSON(&category); err != nil {
		logger.Error("error to unmarshal json", err, zap.String("journey", "createCategory"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("error to unmarshal json"))
		return
	}

	categoryID := uuid.New().String()
	fmt.Println("categoryID", categoryID)

	domain := domain.NewCategoryDomain(
		categoryID,
		category.Title,
		category.Description,
		category.OwnerID,
	)

	fmt.Println("controller", domain)
	if err := cc.service.CreateCategory(domain); err != nil {
		logger.Error("error to create category", err, zap.String("journey", "createCategory"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("error to create category"))
		return
	}

	logger.Info("category created sucesfully", zap.String("journey", "createCategory"))
	c.JSON(http.StatusCreated, "category created sucessfully")
}
