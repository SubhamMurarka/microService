package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresPassword string
	PostgresDatabase string
	PostgresUser     string
	KafkaHost        string
	KafkaPort        string
	KafkaTopic       string
}

var Config AppConfig

func init() {
	err := godotenv.Load("/home/murarka/microService/Payment/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config = AppConfig{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		KafkaHost:        os.Getenv("KAFKA_HOST"),
		KafkaPort:        os.Getenv("KAFKA_PORT"),
		KafkaTopic:       os.Getenv("KAFKA_TOPIC"),
	}
}

// if no value is provided will return fallback value for that variable

// func getEnv(key string, fallback string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return fallback
// }
