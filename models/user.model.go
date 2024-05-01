package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 👈 User struct
type User struct {
	Name               string    `json:"name" bson:"name" binding:"required"`
	Email              string    `json:"email" bson:"email" binding:"required"`
	UID                string    `json:"uid" bson:"uid"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}

// 👈 UserResponse struct
type UserResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	UID       string             `json:"uid,omitempty" bson:"uid,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewUser(name string, email string, uid string) *User {
	return &User{
		Name:         name,
		Email:        email,
		UID:     	  uid,
	}
}

func (model *User) CollectionName() string {
	return "users"
}
