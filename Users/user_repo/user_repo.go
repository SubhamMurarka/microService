package user_repo

import (
	"context"
	"database/sql"

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
	FindUserByName(ctx context.Context, username string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (bool, error)
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
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given email
		}
		return nil, err
	}
	return u, nil
}

func (r *repository) FindUserByName(ctx context.Context, username string) (bool, error) {
	var userID int
	query := "SELECT id FROM users WHERE username = ?"
	err := r.db.QueryRowContext(ctx, query, username).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (bool, error) {
	var userID string
	query := "SELECT id FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
