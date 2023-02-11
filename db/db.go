package db

import (
	"context"
	"fmt"

	"tf2-rcon/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Player struct {
// 	steamID   string
// 	encounter int32
// }

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

func DBAddPlayer(client *mongo.Client, playerID string) {
	collection := client.Database("TF2").Collection("Players")

	filter := bson.D{{Key: "SteamID", Value: playerID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Encounter", Value: 1}}}}
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents updated: %v\n", result.ModifiedCount)
	fmt.Printf("Number of documents upserted: %v\n", result.UpsertedCount)
}
