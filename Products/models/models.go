package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	// Description string `json:"description" bson:"description"`
	Price string `json:"price" bson:"price"`
}

type CreateProduct struct {
	Name string `json:"name" bson:"name"`
	// Description string `json:"description" bson:"description"`
	Price string `json:"price" bson:"price"`
}

type PurchaseReq struct {
	ProductID []string `json:"product_id"`
}

type KafkaEvent struct {
	UserID    string   `json:"user_id"`
	ProductID []string `json:"product_id"`
}

type TokenCreateParams struct {
	Username string
	Email    string
	UserID   string
	jwt.StandardClaims
}
