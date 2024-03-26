package main

import (
	"log"

	"github.com/SubhamMurarka/microService/Users/db"
	"github.com/SubhamMurarka/microService/Users/user_handler"
	"github.com/SubhamMurarka/microService/Users/user_repo"
	"github.com/SubhamMurarka/microService/Users/user_service"
	"github.com/gofiber/fiber/v2"
	"honnef.co/go/tools/config"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialiaze database connection: %s", err)
	}
	Rep := user_repo.NewRepository(dbConn.GetDB())
	userSvc := user_service.NewService(Rep)
	userHandler := user_handler.NewHandler(userSvc)

	app := fiber.New()
	app.Post("/signup", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)

	port := ":" + config.Config.ServerPort

	if err = app.Listen(port); err != nil {
		panic(err)
	}
}
