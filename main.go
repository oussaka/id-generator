package main

//import (
//	"io"
//	"log"
//	"net/http"
//)
import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("Starting Ping the MongoDB database ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("ping mongodb error :%v", err)
		return
	}
	fmt.Println("ping success")
}
