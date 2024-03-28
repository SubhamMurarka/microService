package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name" bson:"name"`
	// Description string `json:"description" bson:"description"`
	Price string `json:"price" bson:"price"`
}

type PurchaseReq struct {
	ID     uuid.UUID `json:"id"`
	UserID string    `json:"userid"`
}
