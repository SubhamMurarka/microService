package models

type Payment struct {
	ProductID []string `json:"product_id"`
	UserID    string   `json:"user_id"`
}
