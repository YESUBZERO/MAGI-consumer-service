package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
	"github.com/YESUBZERO/consumer-service/internal/config"
)

// KafkaHandler maneja los mensajes recibidos
type KafkaHandler struct {
	MessageHandler func(message []byte) error
}

// Iniciar el consumidor Kafka
func StartKafkaConsumer(ctx context.Context, config config.KafkaConfig, messageHandler func(message []byte) error) error {
	consumerGroup, err := sarama.NewConsumerGroup(config.Brokers, config.GroupID, nil)
	if err != nil {
		return err
	}
	defer consumerGroup.Close()

	handler := &KafkaHandler{MessageHandler: messageHandler}

	// Ciclo principal del consumidor
	go func() {
		for {
			if ctx.Err() != nil {
				return
			}
			if err := consumerGroup.Consume(ctx, []string{config.RawTopic}, handler); err != nil {
				log.Printf("Error consumiendo mensajes: %v", err)
			}
		}
	}()

	log.Printf("Consumidor Kafka iniciado para el tópico: %s", config.RawTopic)
	<-ctx.Done()
	return nil
}

func (h *KafkaHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *KafkaHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *KafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := h.MessageHandler(message.Value); err != nil {
			log.Printf("Error procesando mensaje: %v", err)
		}
		session.MarkMessage(message, "")
	}
	return nil
}
