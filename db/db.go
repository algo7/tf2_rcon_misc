package db

import (
	"context"
	"fmt"
	"time"
	"os"

	"tf2-rcon/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect to the database
func Connect() *mongo.Client {
	fmt.Println("Connecting to MongoDB...")

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	mongoURI := "mongodb://localhost:27017"

    if envMongoURI := os.Getenv("MONGODB_URI"); envMongoURI != "" {
        mongoURI = envMongoURI
    }
	
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		utils.ErrorHandler(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		utils.ErrorHandler(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		utils.ErrorHandler(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// AddPlayer adds a player to the database
func AddPlayer(client *mongo.Client, playerID int64, playerName string) *mongo.UpdateResult {
	mongoDBName := "TF2"

    if envMongoDBName := os.Getenv("MONGODB_NAME"); envMongoDBName != "" {
        mongoDBName = envMongoDBName
    }
	
	// Get a handle for your collection
	collection := client.Database(mongoDBName).Collection("Players")

	// Filter by the steamID (64)
	filter := bson.D{{Key: "SteamID", Value: playerID}}

	// The information to be updated
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "SteamID", Value: playerID},
		{Key: "Name", Value: playerName},
		{Key: "UpdatedAt", Value: time.Now().UnixNano()},
	}}}

	// Upsert the document if it doesn't exist
	opts := options.Update().SetUpsert(true)

	// Update the document
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		utils.ErrorHandler(err)
	}

	// fmt.Printf("Number of documents upserted: %v\n", result)
	return result

}
