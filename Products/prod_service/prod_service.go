package prod_service

import (
	"context"
	"fmt"

	"github.com/SubhamMurarka/microService/Products/kafka"
	"github.com/SubhamMurarka/microService/Products/models"
	"github.com/SubhamMurarka/microService/Products/prod_repo"
	"github.com/SubhamMurarka/microService/Products/utils"
)

type service struct {
	prod_repo.Repository
}

type Service interface {
	CreateProduct(ctx context.Context, CreateProduct *models.CreateProduct) (*models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, id string, product *models.CreateProduct) error
	DeleteProduct(ctx context.Context, id string) error
	GetAllProducts(ctx context.Context, page int, pageSize int) ([]models.Product, error)
	Purchase(ctx context.Context, req *models.KafkaEvent) (string, error)
}

func NewService(repository prod_repo.Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateProduct(ctx context.Context, CreateProduct *models.CreateProduct) (*models.Product, error) {
	if CreateProduct.Name == "" || CreateProduct.Price == "" {
		return nil, utils.ErrEmptyFields
	}
	Product, err := s.Repository.CreateProduct(ctx, CreateProduct)
	if err != nil {
		fmt.Println("Cannot create product: ", err)
		return nil, err
	}
	return Product, nil
}

func (s *service) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	if id == "" {
		return nil, utils.ErrEmptyProductID
	}
	product, err := s.Repository.GetProduct(ctx, id)
	if err != nil {
		fmt.Println("Cannot fetch product: ", err)
		return nil, err
	}
	return product, nil
}

func (s *service) UpdateProduct(ctx context.Context, id string, product *models.CreateProduct) error {
	if id == "" || product.Name == "" || product.Price == "" {
		return utils.ErrEmptyFields
	}
	err := s.Repository.UpdateProduct(ctx, id, product)
	if err != nil {
		fmt.Println("Cannot update product: ", err)
	}
	return err
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	if id == "" {
		return utils.ErrEmptyProductID
	}
	err := s.Repository.DeleteProduct(ctx, id)
	if err != nil {
		fmt.Println("Cannot delete product: ", err)
	}
	return err
}

func (s *service) GetAllProducts(ctx context.Context, page int, pageSize int) ([]models.Product, error) {
	products, err := s.Repository.GetAllProducts(ctx, page, pageSize)
	if err != nil {
		fmt.Println("Error fetching all products: ", err)
		return nil, err
	}
	return products, nil
}

func (s *service) Purchase(ctx context.Context, req *models.KafkaEvent) (string, error) {
	if req.UserID == "" || len(req.ProductID) == 0 {
		return "", utils.ErrPurchaseReq
	}

	err := kafka.InitProducer()
	if err != nil {
		fmt.Println("Error initializing producer: ", err)
		return "", err
	}
	defer kafka.CloseKafka()

	err = kafka.PublishMessage(req)
	if err != nil {
		fmt.Println("Error publishing message: ", err)
		return "", err
	}

	return "Payment successfully done", nil
}
