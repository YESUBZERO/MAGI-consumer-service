package config

import "github.com/kelseyhightower/envconfig"

// Configuracion principal del servicio
type Config struct {
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Brokers []string `envconfig:"KAFKA_BROKERS" required:"true"`
	Topic   string   `envconfig:"KAFKA_TOPIC" required:"true"`
	GroupID string   `envconfig:"KAFKA_GROUP_ID" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
