package user_service

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/SubhamMurarka/microService/Users/models"
	"github.com/SubhamMurarka/microService/Users/user_repo"
	"github.com/SubhamMurarka/microService/Users/util"
)

type service struct {
	user_repo.Repository
}

type Service interface {
	CreateUser(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error)
	Login(c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error)
}

func NewService(repository user_repo.Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateUser(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error) {

	if req.Email == "" || req.Username == "" {
		return nil, util.ErrEmptyFields
	}

	u, err := s.Repository.GetUserByEmail(c, req.Email)

	if u != nil {
		return nil, util.ErrEmailExists
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(c, user)
	if err != nil {
		return nil, err
	}

	res := &models.CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}

func (s *service) Login(c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error) {
	u, err := s.Repository.GetUserByEmail(c, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.ErrInvalidCredentials
		}
		return nil, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return nil, util.ErrInvalidCredentials
	}

	token, err := util.GenerateAllTokens(strconv.FormatInt(u.ID, 10), u.Username, u.Email)
	if err != nil {
		return nil, err
	}

	logRes := &models.LoginUserRes{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		Token:    token,
	}

	return logRes, nil
}
