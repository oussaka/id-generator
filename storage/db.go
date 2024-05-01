package storage

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func InitDB() {
	
	fmt.Println("Starting Ping the MongoDB database ...")
	
		ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
		defer cancel()

		var err error
		client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}
	
		/* defer func() {
			fmt.Print("disctonnecte ---------------------")
			if err = clientMongo.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()*/
	
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			log.Fatalf("ping mongodb error :%v", err);
			return
		}

	fmt.Println("Successfully connected to database")
}

func GetClient() *mongo.Client {
	return client
}