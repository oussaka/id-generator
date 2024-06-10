package repositories

import (
	"context"
	"errors"
	"id-generator/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
			UID:       doc.UID,
			Name:      doc.Name,
			Email:     doc.Email,
			CreatedAt: doc.CreatedAt,
			UpdatedAt: doc.UpdatedAt,
		})
	}

	log.Printf("%v documents fetched ", len(queryResult))

	return users, err
}

func GetUser(uid string) (models.User, error) {
	fmt.Println("GET user...")

	var user models.User
	var filter = bson.D{{"uid", uid}}

	err := storage.Collections.Users.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, errors.New("cannot find user")
	}

	log.Printf("%v documents fetched ", user)

	return user, err
}

func FindUserBy(filter map[string]interface{}) (models.User, error) {
	fmt.Println("GET user...", filter)

	var user models.User

	err := storage.Collections.Users.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, err
	}

	if err != nil {
		return user, errors.New("Query return error: " + err.Error())
	}

	log.Printf("%v documents fetched ", user)

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")

	user.CreatedAt = time.Now()
	if user.UID == "" {
		user.UID = services.GenerateUid()
	}

	fmt.Printf("======= USER ================ %#v", user)

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
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	res := storage.Collections.Users.FindOneAndUpdate(context.Background(), filter, updateSet, opts)

	var updatedData models.User

	if err := res.Decode(&updatedData); err != nil {
		log.Fatal(err)
	}

	fmt.Println("======= updated id ================")
	log.Printf("UPDATED ID is : %#v", updatedData.UID)

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
