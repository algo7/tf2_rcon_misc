package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect to the DB
var client = connect()

// connect to the database
func connect() *mongo.Client {
	log.Println("Connecting to MongoDB...")

	// Get the MongoDB URI from the environment
	mongoURI := os.Getenv("MONGODB_URI")

	// If the URI is empty, use the default
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Unable to connect to MongoDB: %v", err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		utils.ErrorHandler(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Unable to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	return client
}
