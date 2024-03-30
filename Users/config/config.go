package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MysqlHost      string
	MysqlPort      string
	MysqlPassword  string
	MysqlDatabase  string
	JwtSecret      string
	ServerPortUser string
}

var Config AppConfig

func init() {
	err := godotenv.Load("/home/murarka/microService/Users/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = AppConfig{
		JwtSecret:      os.Getenv("JWT_SECRET"),
		ServerPortUser: os.Getenv("SERVER_PORT_USER"),
		MysqlHost:      os.Getenv("MYSQL_HOST"),
		MysqlPort:      os.Getenv("MYSQL_PORT"),
		MysqlPassword:  os.Getenv("MYSQL_PASSWORD"),
		MysqlDatabase:  os.Getenv("MYSQL_DATABASE"),
	}
}

// if no value is provided will return fallback value for that variable

// func getEnv(key string, fallback string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return fallback
// }
