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

	// Iniciar el consumidor de Kafka
	err = kafka.StartKafkaConsumer(ctx, kafka.KafkaConfig(cfg.Kafka), processor.ProcessMessage)
	if err != nil {
		log.Fatalf("Error iniciando consumidor de Kafka: %v", err)
	}
}
