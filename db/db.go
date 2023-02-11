package db

import (

	// "go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"

	"tf2-rcon/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	_, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		utils.ErrorHandler(err)
	}

	fmt.Println("Connected to MongoDB!")

	// collection := client.Database("TF2").Collection("Players")
}
