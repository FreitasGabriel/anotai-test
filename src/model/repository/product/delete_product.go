package repository

import (
	"fmt"
	"os"

	"github.com/FreitasGabriel/anotai-test/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (pr *productRepositoryInterface) DeleteProduct(id string) (string, error) {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := pr.databaseConn.Collection(collection_name)
	filter := bson.D{{Key: "id", Value: id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error("error to delete document", err, zap.String("journey", "deleteProduct"))
		return "", err
	}

	fmt.Println("result", result)

	return "product deleted successfully", nil
}
