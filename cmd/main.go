package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YESUBZERO/consumer-service/internal/config"
	"github.com/YESUBZERO/consumer-service/internal/kafka"
	"github.com/YESUBZERO/consumer-service/internal/processor"
)

func main() {
	// Leer configuracion del servicio
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando configuracion: %v", err)
	}

	// Crear contexto para manejar senales de interrupcion
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capturar senales del sistema (Ctrl+C, etc.)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		<-signalChan
		log.Println("Cerrando el servicio...")
		cancel()
	}()

	// Crear el topico de mensajes procesados
	err = kafka.EnsureTopicExists(cfg.Kafka.Brokers, "processed-message", 3, 1)
	if err != nil {
		log.Fatalf("Error creando el topico: %v", err)
	}

	// Iniciar el productor de Kafka
	producer, err := kafka.NewKafkaProducer(cfg.Kafka.Brokers, "processed-messages")
	if err != nil {
		log.Fatalf("Error iniciando el productor Kafka: %v", err)
	}
	defer producer.Producer.Close()

	// Iniciar el consumidor de Kafka
	err = kafka.StartKafkaConsumer(ctx, kafka.KafkaConfig(cfg.Kafka), func(message []byte) error {
		// Procesar y publicar el mensaje
		return processor.ProcessMessage(&kafka.KafkaProducer{Topic: "processed", Producer: nil}, message)
	})
	if err != nil {
		log.Fatalf("Error iniciando consumidor de Kafka: %v", err)
	}
}
