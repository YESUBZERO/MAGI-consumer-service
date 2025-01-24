package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// KafkaHandler es un manejador de mensajes de Kafka
type KafkaHandler struct {
	MessageHandler func(message []byte) error
}

// Configuracion de Kafka
type KafkaConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

// Setup es llamado al inicio de un nuevo consumidor de Kafka
func (handler *KafkaHandler) Setup(sarama.ConsumerGroupSession) error { return nil }

// Cleanup es llamado al finalizar un consumidor de Kafka
func (handler *KafkaHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim es llamado por el consumidor de Kafka para procesar mensajes
func (h *KafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := h.MessageHandler(message.Value); err != nil {
			log.Printf("Error procesando mensaje: %v", err)
		} else {
			session.MarkMessage(message, "")
		}
	}
	return nil
}

// Iniciar el consumidor de Kafka
func StartKafkaConsumer(ctx context.Context, config KafkaConfig, messageHandler func(messagee []byte) error) error {
	// Crear un grupo de consumidores
	consumerGroup, err := sarama.NewConsumerGroup(config.Brokers, config.GroupID, nil)
	if err != nil {
		return err
	}
	defer consumerGroup.Close()

	// Handler para procesar los mensajes
	handler := &KafkaHandler{MessageHandler: messageHandler}

	// Ciclo principal del consumidor
	go func() {
		for {
			if ctx.Err() != nil {
				return
			}

			// Consumir mensajes del Topico
			err := consumerGroup.Consume(ctx, []string{config.Topic}, handler)
			if err != nil {
				log.Printf("Error consumiendo mensajes: %v", err)
			}
		}
	}()

	log.Printf("Consumidor Kafka iniciado para el topico %s", config.Topic)
	<-ctx.Done() // Esperar a que el contexto sea cancelado
	return nil
}
