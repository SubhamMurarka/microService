package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/SubhamMurarka/microService/Products/config"
	"github.com/SubhamMurarka/microService/Products/models"
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

var producer sarama.SyncProducer

func InitProducer() error {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // wait for all Partitions
	// config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	var err error
	producer, err = sarama.NewSyncProducer([]string{kafkaConfig.Host + ":" + kafkaConfig.Port}, config)
	if err != nil {
		return err
	}
	return nil
}

func PublishMessage(message *models.PurchaseReq) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return err
	}

	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: kafkaConfig.Topic,
		Value: sarama.ByteEncoder(messageBytes),
	})

	if err != nil {
		log.Println("Error publishing message to Kafka:", err)
		return err
	}

	fmt.Printf("data inserted in partition %d and with offset value %d", partition, offset)

	return nil
}

func CloseKafka() {
	var err error
	if producer != nil {
		err = producer.Close()
	}
	if err != nil {
		fmt.Println("Kafka not closed", err)
	}
}
