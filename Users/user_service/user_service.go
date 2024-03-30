package user_service

import (
	"context"
	"errors"
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
		return nil, errors.New("email and username cannot be empty")
	}

	exists, err := s.Repository.FindUserByEmail(c, req.Email)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email already exists")
	}

	exists, err = s.Repository.FindUserByName(c, req.Username)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(c, u)

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
		return &models.LoginUserRes{}, errors.New("invalid credentials")
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &models.LoginUserRes{}, errors.New("invalid credentials")
	}

	token, err := util.GenerateAllTokens(strconv.FormatInt(u.ID, 10), u.Username, u.Email)

	if err != nil {
		return &models.LoginUserRes{}, errors.New("try to login again")
	}

	logRes := &models.LoginUserRes{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		Token:    token,
	}

	return logRes, nil
}
