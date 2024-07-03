package database

import (
	"context"
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL  = "MONGODB_URL"
	MONGODB_NAME = "DB_NAME"
)

func NewDatabaseConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_name := os.Getenv(MONGODB_NAME)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("connection successful with database")

	return client.Database(mongodb_name), nil
}
