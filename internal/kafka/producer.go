package kafka

import (
	"log"

	"github.com/IBM/sarama"
	"github.com/YESUBZERO/consumer-service/internal/config"
)

// Producer representa un productor de Kafka
type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(cfg *config.Config) *Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(cfg.Kafka.Brokers, config)
	if err != nil {
		log.Fatalf("Error creando productor de Kafka: %v", err)
	}

	return &Producer{producer: producer}
}

func (p *Producer) SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Error enviando mensaje a Kafka: %v", err)
	}
	//log.Printf("Mensaje enviado a %s: %s", topic, message)
}
