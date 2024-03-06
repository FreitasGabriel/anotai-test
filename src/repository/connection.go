package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/bemobi/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnectionDB(ctx context.Context, Logger *log.Context) *mongo.Client {

	mongoURI := "mongodb://localhost:27017"
	opts := options.Client().SetTimeout(5 * time.Second).ApplyURI(mongoURI)
	client, errDB := mongo.Connect(ctx, opts)
	if errDB != nil {
		Logger.E("Error to connect to database")
	}

	// Check the connection.
	errPing := client.Ping(ctx, nil)
	if errPing != nil {
		fmt.Println("Error to mongoDB!!!", errPing)
	} else {
		Logger.I("Connected to mongoDB!!")
	}

	return client
}
