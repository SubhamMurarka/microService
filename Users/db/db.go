package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SubhamMurarka/microService/Users/config"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string
	Port     string
	Password string
	Database string
}

type Database struct {
	DB *sql.DB
}

var Cfg MysqlConfig

func init() {
	Cfg = MysqlConfig{
		Host:     config.Config.MysqlHost,
		Port:     config.Config.MysqlPort,
		Password: config.Config.MysqlPassword,
		Database: config.Config.MysqlDatabase,
	}
}

func NewDatabase() (*Database, error) {
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s", Cfg.Password, Cfg.Host, Cfg.Port, Cfg.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("not connecting to database : ", err)
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
