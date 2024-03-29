package main

import (
	"fmt"
	"log"

	"github.com/SubhamMurarka/microService/Users/config"
	"github.com/SubhamMurarka/microService/Users/db"
	"github.com/SubhamMurarka/microService/Users/user_handler"
	"github.com/SubhamMurarka/microService/Users/user_repo"
	"github.com/SubhamMurarka/microService/Users/user_service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialiaze database connection: %s", err)
	}

	defer dbConn.Close()

	runDBMigration()

	Rep := user_repo.NewRepository(dbConn.GetDB())
	userSvc := user_service.NewService(Rep)
	userHandler := user_handler.NewHandler(userSvc)

	app := fiber.New()
	app.Post("/signup", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)
	app.Get("/auth", user_handler.Authorise)
	port := ":" + config.Config.ServerPortUser

	if err = app.Listen(port); err != nil {
		panic(err)
	}
}

func runDBMigration() {
	dsn := fmt.Sprintf("mysql://root:%s@tcp(%s:%s)/%s", config.Config.MysqlPassword, config.Config.MysqlHost, config.Config.MysqlPort, config.Config.MysqlDatabase)
	migrationsPath := "file://db/migrations"

	m, err := migrate.New(
		dsn,
		migrationsPath,
	)
	if err != nil {
		log.Fatal("cannot create new migrate instance", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("unable to migrate up ", err)
	}

	log.Println("DB migrated successfully")
}
