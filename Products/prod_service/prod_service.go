package prod_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SubhamMurarka/microService/Products/kafka"
	"github.com/SubhamMurarka/microService/Products/models"
	"github.com/SubhamMurarka/microService/Products/prod_repo"
)

type service struct {
	prod_repo.Repository
	timeout time.Duration
}

type Service interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, id string, product *models.Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetAllProducts(ctx context.Context, page int) ([]*models.Product, error)
	Purchase(ctx context.Context, req *models.PurchaseReq) (string, error)
}

func NewService(repository prod_repo.Repository) Service {
	return &service{
		repository,
		time.Duration(8) * time.Second,
	}
}

const defualtPageSize = 10

func (s *service) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	if product == nil {
		return nil, errors.New("invalid product")
	}
	return s.Repository.CreateProduct(ctx, product)
}

func (s *service) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	if id == "" {
		return nil, errors.New("empty product id")
	}
	return s.Repository.GetProduct(ctx, id)
}

func (s *service) UpdateProduct(ctx context.Context, id string, product *models.Product) error {
	if id == "" || product == nil {
		return errors.New("invalid product ID or nil product")
	}
	return s.Repository.UpdateProduct(ctx, id, product)
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("empty product ID")
	}
	return s.Repository.DeleteProduct(ctx, id)
}

func (s *service) GetAllProducts(ctx context.Context, page int) ([]*models.Product, error) {
	if page < 1 {
		return nil, errors.New("invalid page or pageSize")
	}
	pageSize := defualtPageSize
	return s.Repository.GetAllProducts(ctx, page, pageSize)
}

func (s *service) Purchase(ctx context.Context, req *models.PurchaseReq) (string, error) {
	err := kafka.InitProducer()

	defer kafka.CloseKafka()

	if err != nil {
		fmt.Println("error initialising producer", err)
		return "cannot process the Payment", err
	}

	err = kafka.PublishMessage(req)
	if err != nil {
		fmt.Println("error in publishing message", err)
		return "cannot process the Payment", err
	}

	return "Payment successfully Done", nil
}
