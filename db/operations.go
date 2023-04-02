package db

import (
	"context"
	"os"
	"tf2-rcon/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Database  name
	mongoDBName = os.Getenv("MONGODB_NAME")
)

// AddPlayer adds a player to the database
func AddPlayer(playerID int64, playerName string) *mongo.UpdateResult {

	// If the URI is empty, use the default
	if mongoDBName == "" {
		mongoDBName = "TF2"
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
		utils.ErrorHandler(err, false)
	}

	// fmt.Printf("Number of documents upserted: %v\n", result)
	return result

}

// AddPlayer adds the given chat message to the database
func AddChat(playerID int64, playerName string) *mongo.UpdateResult {

	// If the URI is empty, use the default
	if mongoDBName == "" {
		mongoDBName = "TF2"
	}

	// Get a handle for your collection
	collection := client.Database(mongoDBName).Collection("Chats")

	// Filter by the steamID (64)
	filter := bson.D{{Key: "SteamID", Value: playerID}}

	// The information to be updated
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "SteamID", Value: playerID},
		{Key: "Name", Value: playerName},
		{Key: "UpdatedAt", Value: time.Now().UnixNano()},
	}}}

	// Upsert the document if it doesn't exist
	opts := options.InsertOne()

	// Update the document
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		utils.ErrorHandler(err, false)
	}

	// fmt.Printf("Number of documents upserted: %v\n", result)
	return result

}
