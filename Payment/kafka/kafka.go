package kafka

import (
	"fmt"

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
	conn, err := sarama.NewConsumer(brokerUrl, config)
	if err != nil {
		fmt.Println("not able to create new consumer : ", err)
		return nil, err
	}

	return conn, nil
}
