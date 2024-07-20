package main

import (
	"context"
	"log"

	"github.com/FreitasGabriel/anotai-test/src/configuration/database"
	configuration "github.com/FreitasGabriel/anotai-test/src/configuration/dependencies"
	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"github.com/FreitasGabriel/anotai-test/src/configuration/rest_err"
	"github.com/FreitasGabriel/anotai-test/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("starting application")
	err := godotenv.Load()
	if err != nil {
		logger.Error(
			"error to initialize environment variables",
			rest_err.NewInternalServerError("error to initialize environment variables"),
		)
		return
	}

	db, err := database.NewDatabaseConnection(context.Background())
	if err != nil {
		logger.Error("error to connect database", err)
		return
	}

	productController, categoryController := configuration.InitDep(db)

	router := gin.Default()
	routes.InitProductRoutes(&router.RouterGroup, productController)
	routes.InitCategoryRoutes(&router.RouterGroup, categoryController)

	if err := router.Run(":8000"); err != nil {
		log.Fatal("error to start server")
	}

	logger.Info("application running")
}
