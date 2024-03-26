package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"honnef.co/go/tools/config"
)

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Database struct {
	DB *sql.DB
}

var cfg MysqlConfig

func init() {
	cfg = MysqlConfig{
		Host:     config.Config.Host,
		Port:     config.Config.Port,
		User:     config.Config.User,
		Password: config.Config.Password,
		Database: config.Config.Database,
	}
}

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dsn)
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
