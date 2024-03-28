package db

import (
	"context"
	"fmt"

	"github.com/SubhamMurarka/microService/Products/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string
	Password string
	Port     string
	User     string
	Database string
}

type MongoDB struct {
	client *mongo.Client
	DB     *mongo.Database
}

var cfg MongoConfig

func init() {
	cfg = MongoConfig{
		Host:     config.Config.PostgresHost,
		Password: config.Config.PostgresPassword,
		Port:     config.Config.PostgresPort,
		User:     config.Config.PostgresUser,
		Database: config.Config.PostgresDatabase,
	}
}

func NewDatabase() (*MongoDB, error) {
	URI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	db := client.Database("product")

	return &MongoDB{
		client: client,
		DB:     db,
	}, nil
}
