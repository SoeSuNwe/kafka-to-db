package main

import (
	"fmt"
	"kafka-to-db/db"
	"kafka-to-db/kafka"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read DB connection details from .env
	dbConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"))

	// Connect to the database
	dbConn, err := db.Connect(dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Read Kafka configuration from .env
	kafkaConfig := kafka.Config{
		Brokers: []string{os.Getenv("KAFKA_BROKERS")},
		Topic:   os.Getenv("KAFKA_TOPIC"),
		GroupID: os.Getenv("KAFKA_GROUP_ID"),
	}

	// Initialize Kafka consumer
	consumer := kafka.NewConsumer(kafkaConfig, dbConn)
	consumer.Consume()
}
