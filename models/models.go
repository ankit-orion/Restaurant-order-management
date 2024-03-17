package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID     primitive.ObjectID `bson:"_id"`
	Dish   string             `json:"dish" validate:"required"` //validator that is required to file in
	Price  float64            `json:"price" `
	Server string             `json:"server" validate:"required"`
	Table  string             `json:"table"`
}

type Waiter struct{
	Server 	string	`json:"server"`
}