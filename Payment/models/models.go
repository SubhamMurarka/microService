package models

type Payment struct {
	ProductID []string `json:"id"`
	UserID    string   `json:"userid"`
}
