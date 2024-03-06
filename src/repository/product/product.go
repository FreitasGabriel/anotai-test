package product

import (
	"context"

	"github.com/FreitasGabriel/anotai-test/src/model"
	"github.com/bemobi/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDB struct {
	DB      *mongo.Client
	Logger  *log.Context
	Context *context.Context
}

func (db *ProductDB) Create(payload *model.ProductPayload) error {

	productPayload := model.ProductPayload{
		OwnerID:     payload.OwnerID,
		ProductID:   payload.ProductID,
		Title:       payload.Title,
		Description: payload.Description,
		Price:       payload.Price,
		Category:    payload.Category,
	}

	_, err := db.DB.Database("anotai-test").Collection("products").InsertOne(*db.Context, productPayload)
	if err != nil {
		db.Logger.E("Error to insert document in colletion", err)
		return err
	}

	return nil
}

func (db *ProductDB) List() ([]model.ProductPayload, error) {

	var products []model.ProductPayload

	cursor, err := db.DB.Database("anotai-test").Collection("products").Find(*db.Context, bson.D{})
	if err != nil {
		db.Logger.E("Error to return all documents from database", err)
		return nil, err
	}

	for cursor.Next(*db.Context) {
		var product model.ProductPayload
		err := cursor.Decode(&product)
		if err != nil {
			db.Logger.E("Error to decoder colletion from db", err)
		}
		products = append(products, product)
	}

	return products, nil
}
