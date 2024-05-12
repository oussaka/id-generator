package repositories

import (
	"context"
	"log"
	"fmt"
	"id-generator/models"
	"id-generator/storage"
)

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")
	insertedResult, err := storage.Collections.Users.InsertOne(context.TODO(), user)

	if err != nil {
		log.Fatalf("inserted error : %v", err)
		return models.User{}, err
	}
	fmt.Println("======= inserted id ================")
	log.Printf("inserted ID is : %v", insertedResult.InsertedID)
	
return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
return user, nil
}
