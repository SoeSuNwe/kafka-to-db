package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaConfig KafkaConfig
	DBConfig    DBConfig
}

type KafkaConfig struct {
	Brokers []string `envconfig:"KAFKA_BROKERS" required:"true"`
	Topic   string   `envconfig:"KAFKA_TOPIC" required:"true"`
	GroupID string   `envconfig:"KAFKA_GROUP_ID" required:"true"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
