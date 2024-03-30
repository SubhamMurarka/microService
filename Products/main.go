package main

import (
	"log"

	"github.com/SubhamMurarka/microService/Products/auth"
	"github.com/SubhamMurarka/microService/Products/config"
	"github.com/SubhamMurarka/microService/Products/db"
	"github.com/SubhamMurarka/microService/Products/prod_handler"
	"github.com/SubhamMurarka/microService/Products/prod_repo"
	"github.com/SubhamMurarka/microService/Products/prod_service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	mongoDB, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	repo := prod_repo.NewMongoRepository(mongoDB.DB, "items")
	prodSvc := prod_service.NewService(repo)
	ProductHandler := prod_handler.NewHandler(prodSvc)

	app := fiber.New()

	app.Use(auth.Authorise)
	app.Post("/products", ProductHandler.CreateProduct)
	app.Get("/products/:id", ProductHandler.GetProduct)
	app.Put("/products/:id", ProductHandler.UpdateProduct)
	app.Delete("/products/:id", ProductHandler.DeleteProduct)
	app.Get("/products", ProductHandler.GetAllProducts)
	app.Post("/products/purchase", ProductHandler.Purchase)

	port := ":" + config.Config.ServerPortProduct

	if err = app.Listen(port); err != nil {
		panic(err)
	}
}
