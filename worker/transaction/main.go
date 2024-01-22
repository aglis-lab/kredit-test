package main

import (
	"context"
	"kreditplus/src/app"
	"kreditplus/src/entity"
	"kreditplus/src/v1/contract"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()

	// Init app context
	if err := app.Init(ctx); err != nil {
		log.Panic(err)
	}

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{app.Config().Kafka.Host},
		Topic:     app.Config().Kafka.TransactionTopic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	// TODO: Calculate Offset

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}

		transaction := entity.Transaction{}
		err = app.GormDB().Where("txn_id", string(message.Value)).First(&transaction).Error
		if err != nil {
			// TODO: Handle Transaction Error for Worker
			log.Println("Transaction Error", err)
		}

		// TODO: Doing very complex logic
		// ex: send a money to partner
		if transaction.Status == contract.TransactionSettlement {
			transaction.Status = contract.TransactionComplete

			err = app.GormDB().Save(&transaction).Error
			if err != nil {
				// TODO: Handle Transaction Error for Worker
				log.Println("Transaction Error", err)
			}
		}

		log.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
