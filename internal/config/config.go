package config

import (
	"log"

	"github.com/YESUBZERO/consumer-service/internal/models"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configuración principal del servicio
type Config struct {
	Kafka KafkaConfig
	DB    DatabaseConfig
}

// Configuración de Kafka
type KafkaConfig struct {
	Brokers []string `envconfig:"KAFKA_BROKERS" required:"true"`
	//DynamicTopic  string   `envconfig:"KAFKA_DYNAMIC_TOPIC" required:"true"`
	StaticTopic   string `envconfig:"KAFKA_STATIC_TOPIC" required:"true"`
	ScrapeTopic   string `envconfig:"KAFKA_SCRAPE_TOPIC" default:""`
	EnrichedTopic string `envconfig:"KAFKA_ENRICHED_TOPIC" default:""`
	GroupID       string `envconfig:"KAFKA_GROUP_ID" required:"true"`
}

type DatabaseConfig struct {
	DSN string `envconfig:"DATABASE_DSN" required:"true"`
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

// Inizializar la configuración de la base de datos
func InitDB(cfg DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	// Auto-migrar el esquema de la base de datos
	err = db.AutoMigrate(&models.Ship{})
	if err != nil {
		log.Fatalf("Error migrando el esquema de la base de datos: %v", err)
	}

	return db
}
