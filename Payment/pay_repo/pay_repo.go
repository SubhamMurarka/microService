package pay_repo

import (
	"context"
	"fmt"

	"github.com/SubhamMurarka/microService/Payment/db"
	"github.com/SubhamMurarka/microService/Payment/models"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}

type Repository interface {
	CreatePayment(ctx context.Context, payment *models.Payment) error
}

func (r *repository) CreatePayment(ctx context.Context, payment *models.Payment) error {
	var lastInsertId int
	query := "INSERT INTO payments(user_id, product_id) VALUES ($1, $2) returning id"
	for _, productID := range payment.ProductID {
		err := r.db.QueryRowContext(ctx, query, payment.UserID, productID).Scan(&lastInsertId)
		if err != nil {
			return err
		}
		fmt.Println(lastInsertId)
	}
	return nil
}
