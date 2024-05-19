package kafka

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/SubhamMurarka/microService/Payment/config"
)

type KafkaConfig struct {
	Host  string
	Port  string
	Topic string
}

var kafkaConfig KafkaConfig

func init() {
	kafkaConfig = KafkaConfig{
		Host:  config.Config.KafkaHost,
		Port:  config.Config.KafkaPort,
		Topic: config.Config.KafkaTopic,
	}
}

var brokerUrl []string

func ConnectConsumer() (sarama.Consumer, error) {
	url := fmt.Sprintf("%s:%s", kafkaConfig.Host, kafkaConfig.Port)
	brokerUrl = append(brokerUrl, url)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	var conn sarama.Consumer
	var err error

	for retries := 10; retries > 0; retries-- {
		conn, err = sarama.NewConsumer(brokerUrl, config)
		if err == nil {
			break
		}
		log.Printf("error creating Kafka consumer, retrying: %v", err)
		time.Sleep(10 * time.Second)
	}

	if err != nil {
		log.Printf("not able to create new consumer after retries: %v", err)
		return nil, err
	}

	return conn, nil
}
