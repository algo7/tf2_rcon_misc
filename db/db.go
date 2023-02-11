package db

import (
	"context"
	"fmt"

	"tf2-rcon/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	steamID   string
	encounter int32
}

func DBConnect() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		utils.ErrorHandler(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
