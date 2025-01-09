package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-to-db/db"
	"kafka-to-db/models"
	"log"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	Brokers []string
	Topic   string
	GroupID string
}

type Consumer struct {
	config Config
	db     *db.DB
}

// NewConsumer initializes a new Kafka consumer
func NewConsumer(config Config, db *db.DB) *Consumer {
	return &Consumer{config: config, db: db}
}

// Consume reads messages from Kafka and upserts them into the database
func (c *Consumer) Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.config.Brokers,
		Topic:   c.config.Topic,
		GroupID: c.config.GroupID,
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		fmt.Printf("Received message: %s\n", msg.Value)

		var message models.Message
		if err := json.Unmarshal(msg.Value, &message); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		if err := c.db.SaveMessage(message); err != nil {
			log.Printf("Failed to save message: %v", err)
		}
	}
}
