package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MongoHost         string
	MongoPort         string
	ServerPortProduct string
	ServerPortUser    string
	KafkaHost         string
	KafkaPort         string
	KafkaTopic        string
}

var Config AppConfig

func init() {
	err := godotenv.Load("/home/murarka/microService/Products/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Config = AppConfig{
		ServerPortProduct: os.Getenv("SERVER_PORT_PRODUCT"),
		MongoHost:         os.Getenv("MONGO_HOST"),
		MongoPort:         os.Getenv("MONGO_PORT"),
		KafkaHost:         os.Getenv("KAFKA_HOST"),
		KafkaPort:         os.Getenv("KAFKA_PORT"),
		KafkaTopic:        os.Getenv("KAFKA_TOPIC"),
		ServerPortUser:    os.Getenv("SERVER_PORT_USER"),
	}
}

// if no value is provided will return fallback value for that variable

// func getEnv(key string, fallback string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return fallback
// }
