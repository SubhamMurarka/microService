package main

import (
	"database/sql"
	"fmt"
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
		log.Fatalf("could not initialiaze database connection: %s", err)
	}

	err = dbConn.DB.Ping()
	if err != nil {
		fmt.Println("connection is not setup ", err)
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
		panic(err)
	}
}

func runDBMigration(db *sql.DB) {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("cannot create new migrate instance", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("unable to migrate up ", err)
	}

	log.Println("DB migrated successfully")
}
