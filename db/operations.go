package db

import (
	"context"
	"os"
	"tf2-rcon/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Database  name
	mongoDBName = os.Getenv("MONGODB_NAME")
)

// Document structs
type Player struct {
	SteamID   int64  `bson:"SteamID"`
	Name      string `bson:"Name"`
	UpdatedAt int64  `bson:"UpdatedAt"`
}

type Chat struct {
	SteamID   int64  `bson:"SteamID"`
	Name      string `bson:"Name"`
	Message   string `bson:"message,omitempty"`
	UpdatedAt int64  `bson:"updatedAt"`
}

// AddPlayer adds a player to the database
func AddPlayer(player Player) *mongo.UpdateResult {

	// If the URI is empty, use the default
	if mongoDBName == "" {
		mongoDBName = "TF2"
	}

	// Get a handle for your collection
	collection := client.Database(mongoDBName).Collection("Players")

	// Filter by the steamID (64)
	filter := bson.D{{Key: "SteamID", Value: player.SteamID}}

	// The information to be updated
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "SteamID", Value: player.SteamID},
		{Key: "Name", Value: player.Name},
		{Key: "UpdatedAt", Value: player.UpdatedAt},
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

// AddChat adds a chat message to the database
func AddChat(chat Chat) *mongo.InsertOneResult {

	// If the URI is empty, use the default
	if mongoDBName == "" {
		mongoDBName = "TF2"
	}

	// Get a handle for your collection
	collection := client.Database(mongoDBName).Collection("Chats")

	// The information to be updated
	insert := bson.D{
		{Key: "SteamID", Value: chat.SteamID},
		{Key: "Name", Value: chat.Name},
		{Key: "Message", Value: chat.Message},
		{Key: "UpdatedAt", Value: chat.UpdatedAt},
	}

	// Update the document
	result, err := collection.InsertOne(context.TODO(), insert)

	if err != nil {
		utils.ErrorHandler(err, false)
	}

	// fmt.Printf("Number of documents upserted: %v\n", result)
	return result
}
