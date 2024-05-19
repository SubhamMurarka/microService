package config

import (
	"os"
)

type AppConfig struct {
	MongoHost         string
	MongoPort         string
	ServerPortProduct string
	PostgresHost      string
	PostgresPort      string
	PostgresPassword  string
	PostgresDatabase  string
	KafkaHost         string
	KafkaPort         string
	KafkaTopic        string
	JwtSecret         string
}

var Config AppConfig

func init() {
	Config = AppConfig{
		MongoHost:         os.Getenv("MONGO_HOST"),
		MongoPort:         os.Getenv("MONGO_PORT"),
		ServerPortProduct: os.Getenv("SERVER_PORT_PRODUCT"),
		PostgresHost:      os.Getenv("POSTGRES_HOST"),
		PostgresPort:      os.Getenv("POSTGRES_PORT"),
		PostgresPassword:  os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase:  os.Getenv("POSTGRES_DATABASE"),
		KafkaHost:         os.Getenv("KAFKA_HOST"),
		KafkaPort:         os.Getenv("KAFKA_PORT"),
		KafkaTopic:        os.Getenv("KAFKA_TOPIC"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
	}
}
