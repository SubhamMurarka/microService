package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialiaze database connection: %s", err)
	}

	defer dbConn.Close()

	runDBMigration()
	repo := pay_repo.NewRepository(dbConn.DB)

	consumer, _ := kafka.ConnectConsumer()
	PartitionConsumer, err := consumer.ConsumePartition(config.Config.KafkaTopic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("PartitionConsumer not created : %s", err)

	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-PartitionConsumer.Errors():
				fmt.Println(err)
			case msg := <-PartitionConsumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
				var payment models.Payment
				if err := json.Unmarshal(msg.Value, &payment); err != nil {
					fmt.Println("Error parsing message:", err)
					continue
				}
				if id, err := repo.CreatePayment(context.TODO(), &payment); err != nil {
					fmt.Println("Error writing to database:", err)
					continue
				} else {
					fmt.Println("payment done with id : ", id)
				}
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	err = consumer.Close()
	if err != nil {
		fmt.Println("error closing the consumer : ", err)
	}
}

func runDBMigration() {
	dsn := fmt.Sprintf("postgres://root:%s@%s:%s/%s?sslmode=disable", config.Config.PostgresPassword, config.Config.PostgresHost, config.Config.PostgresPort, config.Config.PostgresDatabase)
	migrationsPath := "file://db/migration"

	m, err := migrate.New(
		dsn,
		migrationsPath,
	)
	if err != nil {
		log.Fatal("cannot create new migrate instance", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("unable to migrate up ", err)
	}

	log.Println("DB migrated successfully")
}
