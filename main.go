package main

import (
	"fmt"
	"kafka-to-db/config"
	"kafka-to-db/db"
	"kafka-to-db/kafka"
	"log"

	"github.com/joho/godotenv"
)

func main() {
// Load .env file
if err := godotenv.Load(); err != nil {
	log.Fatalf("Error loading .env file: %v", err)
}

// Load configuration from environment variables
cfg, err := config.LoadConfig()
if err != nil {
	log.Fatalf("Failed to load config: %v", err)
}

// Print out the loaded configuration (for debugging)
fmt.Printf("Loaded config: %+v\n", cfg)

// Use the loaded Kafka and DB configurations
dbConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	cfg.DBConfig.Host,
	cfg.DBConfig.User,
	cfg.DBConfig.Password,
	cfg.DBConfig.Name,
	cfg.DBConfig.Port)

// Connect to the database
dbConn, err := db.Connect(dbConnStr)
if err != nil {
	log.Fatalf("Failed to connect to the database: %v", err)
}

// Initialize Kafka consumer
kafkaConfig := kafka.Config{
	Brokers: cfg.KafkaConfig.Brokers,
	Topic:   cfg.KafkaConfig.Topic,
	GroupID: cfg.KafkaConfig.GroupID,
}

// Initialize Kafka consumer and start consuming
consumer := kafka.NewConsumer(kafkaConfig, dbConn)
consumer.Consume()
}
