package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/SubhamMurarka/microService/Payment/config"
	"github.com/SubhamMurarka/microService/Payment/db"
	"github.com/SubhamMurarka/microService/Payment/kafka"
	"github.com/SubhamMurarka/microService/Payment/models"
	"github.com/SubhamMurarka/microService/Payment/pay_repo"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	defer dbConn.Close()

	runDBMigration(dbConn.DB)
	repo := pay_repo.NewRepository(dbConn.DB)

	consumer, err := kafka.ConnectConsumer()
	if err != nil {
		log.Fatalf("failed to connect to Kafka: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("error closing the consumer: %v", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(config.Config.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("failed to start partition consumer: %v", err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				log.Printf("error from partition consumer: %v", err)
			case msg := <-partitionConsumer.Messages():
				msgCount++
				log.Printf("received message count %d: | Topic(%s) | Message(%s)\n", msgCount, string(msg.Topic), string(msg.Value))
				var payment models.Payment
				if err := json.Unmarshal(msg.Value, &payment); err != nil {
					log.Printf("error parsing message: %v", err)
					continue
				}
				if err := repo.CreatePayment(context.Background(), &payment); err != nil {
					log.Printf("error writing to database: %v", err)
					continue
				}
			case <-sigchan:
				log.Println("interrupt detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	log.Printf("processed %d messages", msgCount)
}

func runDBMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("error creating instance: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"postgres", driver)
	if err != nil {
		log.Fatalf("cannot create new migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("unable to migrate up: %v", err)
	}

	log.Println("DB migrated successfully")
}
