package config

import (
	"os"
)

type AppConfig struct {
	PostgresHost      string
	PostgresPort      string
	PostgresUser      string
	PostgresPassword  string
	PostgresDatabase  string
	ServerPortProduct string
	ServerPortUser    string
	KafkaHost         string
	KafkaPort         string
	KafkaTopic        string
}

var Config AppConfig

func init() {
	Config = AppConfig{
		ServerPortProduct: os.Getenv("SERVER_PORT_PRODUCT"),
		PostgresHost:      os.Getenv("POSTGRES_HOST"),
		PostgresPort:      os.Getenv("POSTGRES_PORT"),
		PostgresUser:      os.Getenv("POSTGRES_USER"),
		PostgresPassword:  os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase:  os.Getenv("POSTGRES_DATABASE"),
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
