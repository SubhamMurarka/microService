package user_repo

import (
	"context"

	"github.com/SubhamMurarka/microService/Users/db"
	"github.com/SubhamMurarka/microService/Users/models"
)

type repository struct {
	db db.DBTX
}

func NewRepository(db db.DBTX) Repository {
	return &repository{db: db}
}

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := "INSERT INTO users(username, email, password) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = lastInsertId
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, username, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}
