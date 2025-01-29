package main

import (
	"log"

	"github.com/YESUBZERO/consumer-service/internal/config"
	"github.com/YESUBZERO/consumer-service/internal/kafka"
	"github.com/YESUBZERO/consumer-service/internal/repository"
)

func main() {
	// Cargar configuración desde variables de entorno
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Inicializar conexión a PostgreSQL
	db := config.InitDB(cfg.DB)

	// Inicializar repositorio para PostgreSQL
	shipRepo := repository.NewShipRepository(db)

	// Inicializar productor de Kafka
	producer := kafka.NewProducer(cfg)

	// Inicializar consumidor de Kafka
	consumer := kafka.NewConsumer(cfg, producer, shipRepo)

	// Iniciar el consumidor en una goroutine
	go consumer.ConsumeMessages()

	// Mantener el servicio corriendo
	log.Println("📡 Storage Service en ejecución...")
	select {}
}
