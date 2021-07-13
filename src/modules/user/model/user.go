package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"password" bson:"password"`
}

type UserForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type DeleteUser struct {
	Id primitive.ObjectID `json:"id"`
}
type LoginResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
