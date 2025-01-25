package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuración principal del servicio
type Config struct {
	Kafka KafkaConfig
}

// Configuración de Kafka
type KafkaConfig struct {
	Brokers        []string `envconfig:"KAFKA_BROKERS" required:"true"`
	RawTopic       string   `envconfig:"KAFKA_RAW_TOPIC" required:"true"`
	ProcessedTopic string   `envconfig:"KAFKA_PROCESSED_TOPIC" required:"true"`
	GroupID        string   `envconfig:"KAFKA_GROUP_ID" required:"true"`
}

// Cargar configuración desde variables de entorno
func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
