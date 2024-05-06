package main

import (
	"database/sql"
	"log"

	"github.com/SubhamMurarka/microService/Users/config"
	"github.com/SubhamMurarka/microService/Users/db"
	"github.com/SubhamMurarka/microService/Users/user_handler"
	"github.com/SubhamMurarka/microService/Users/user_repo"
	"github.com/SubhamMurarka/microService/Users/user_service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	if err = dbConn.DB.Ping(); err != nil {
		log.Fatalf("Database connection is not set up: %s", err)
	}

	defer dbConn.Close()

	runDBMigration(dbConn.DB)

	Rep := user_repo.NewRepository(dbConn.DB)
	userSvc := user_service.NewService(Rep)
	userHandler := user_handler.NewHandler(userSvc)

	app := fiber.New()
	app.Post("/signup", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)
	app.Get("/auth", user_handler.Authorise)
	port := ":" + config.Config.ServerPortUser

	if err = app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func runDBMigration(db *sql.DB) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("Cannot create migrate instance: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply database migration: %s", err)
	}

	log.Println("DB migrated successfully")
}
