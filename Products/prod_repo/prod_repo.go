package prod_repo

import (
	"context"
	"fmt"
	"log"

	"github.com/SubhamMurarka/microService/Products/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database, collectionName string) Repository {
	return &mongoRepository{
		collection: database.Collection(collectionName),
	}
}

type Repository interface {
	CreateProduct(ctx context.Context, product *models.CreateProduct) (*models.Product, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, id string, product *models.CreateProduct) error
	DeleteProduct(ctx context.Context, id string) error
	GetAllProducts(ctx context.Context, page int, pageSize int) ([]models.Product, error)
}

func (r *mongoRepository) CreateProduct(ctx context.Context, product *models.CreateProduct) (*models.Product, error) {
	inserted, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("product got inserted with inserted id: ", inserted.InsertedID)
	products := models.Product{
		ID:    inserted.InsertedID.(primitive.ObjectID),
		Name:  product.Name,
		Price: product.Price,
	}
	return &products, nil
}

func (r *mongoRepository) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.collection.FindOne(ctx, primitive.M{"_id": _id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *mongoRepository) UpdateProduct(ctx context.Context, id string, product *models.CreateProduct) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{"$set": product}
	result, err := r.collection.UpdateOne(ctx, primitive.M{"_id": _id}, update)
	fmt.Println("product got updated with modify count: ", result.ModifiedCount)
	return err
}

func (r *mongoRepository) DeleteProduct(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.collection.DeleteOne(ctx, primitive.M{"_id": _id})
	fmt.Println("product got deleted with delete count: ", result.DeletedCount)
	return err
}

func (r *mongoRepository) GetAllProducts(ctx context.Context, page int, pageSize int) ([]models.Product, error) {
	var products []models.Product

	skip := (page - 1) * pageSize

	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products, nil
}
