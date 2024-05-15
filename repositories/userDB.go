package repositories

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"

	"fmt"
	//	"github.com/labstack/echo/v4"
	"id-generator/models"
	"id-generator/storage"
)


func GetUsers() ([]models.User, error) {
	fmt.Println("GET users...")

	var opts = options.Find().SetSort(bson.M{
		"created_at": -1,
	})

	cursor, err := storage.Collections.Users.Find(context.Background(), bson.M{}, opts)

	if err != nil {
		log.Fatalf("find collection err : %v", err)
	}

	var queryResult []models.User
	var users []models.User

	if err := cursor.All(context.Background(), &queryResult); err != nil {
		log.Fatalf("find collection err : %v", err)
	}

	for _, doc := range queryResult {
		users = append(users, models.User{
			UID: doc.UID,
			Name: doc.Name,
			Email: doc.Email,
			CreatedAt: doc.CreatedAt,
			UpdatedAt: doc.UpdatedAt,
		})
	}

	log.Printf("%v documents fetched ", len(queryResult))

	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")
	user.CreatedAt = time.Now()
	insertedResult, err := storage.Collections.Users.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatalf("inserted error : %v", err)
		return models.User{}, err
	}
	fmt.Println("======= inserted id ================")
	log.Printf("inserted ID is : %v", insertedResult.InsertedID)

	return user, nil
}

func UpdateUser(user models.User, uid string) (models.User, error) {
	fmt.Println("Updating user...")

	filter := bson.M{"uid": bson.M{"$eq": uid}}
	updateSet := bson.M{"$set": bson.M{"updated_at": time.Now()}}

	result, err := storage.Collections.Users.UpdateOne(context.Background(), filter, updateSet)

	if err != nil {
		log.Fatalf("updated error : %v", err)
		return models.User{}, err
	}

	fmt.Println("======= updated id ================")
	log.Printf("UPDATED COUNT is : %#v", result.ModifiedCount)
	log.Printf("UPDATED ID is : %#v", result)

	opts := options.FindOne()
	err = storage.Collections.Users.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		fmt.Printf("No documents found  %#v", err)

		return user, errors.New(fmt.Sprintf("update user error : No user found with uid %v", uid))
	}

	return user, nil
}

func DeleteUser(user models.User, uid string) (models.User, error) {
	fmt.Println("Deleting user...")

	filter := bson.M{"uid": bson.M{"$eq": uid}}

	result, err := storage.Collections.Users.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatalf("delete user error : %v", err)
		return models.User{}, err
	}

	fmt.Println("======= deleted id ================")
	log.Printf("deleted COUNT is : %#v", result.DeletedCount)

	return user, nil
}
