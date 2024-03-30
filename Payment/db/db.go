package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SubhamMurarka/microService/Payment/config"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Database struct {
	DB *sql.DB
}

var Cfg Config

func init() {
	Cfg = Config{
		Host:     config.Config.PostgresHost,
		Port:     config.Config.PostgresPort,
		Password: config.Config.PostgresPassword,
		Database: config.Config.PostgresDatabase,
	}
}

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf("postgres://postgres:%s@%s:%s/%s?sslmode=disable", Cfg.Password, Cfg.Host, Cfg.Port, Cfg.Database)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func (d *Database) Close() {
	d.DB.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.DB
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
