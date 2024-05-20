package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Collections struct {
	Users *mongo.Collection
}

func InitDB() {
	log.Print("Initializing database connection")

	var err error
	client, err = connectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dbName := os.Getenv("MONGO_DATABASE")
	Collections.Users = client.Database(dbName).Collection("users")
}

func connectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new client: %v", err)
	}

	/* defer func() {
		fmt.Print("disctonnecte ---------------------")
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}() */

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to database")

	return client, nil
}

func GetClient() *mongo.Client {
	return client
}
