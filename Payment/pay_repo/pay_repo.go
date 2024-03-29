package pay_repo

import (
	"context"

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
	CreatePayment(ctx context.Context, payment *models.Payment) (int, error)
}

func (r *repository) CreatePayment(ctx context.Context, payment *models.Payment) (int, error) {
	var lastInsertId int
	query := "INSERT INTO payments(user_id, product_id) VALUES ($1, $2) returning id"
	err := r.db.QueryRowContext(ctx, query, payment.UserID, payment.ProductID).Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}
	lastInsertId = int(lastInsertId)
	return lastInsertId, nil
}
