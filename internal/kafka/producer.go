package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

// KafkaProducer es un productor de mensajes de Kafka
type KafkaProducer struct {
	Producer sarama.SyncProducer
	Topic    string
}

// NewKafkaProducer crea un nuevo productor de Kafka
func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	// Configuracion del productor
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll

	// Crear el productor
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Error creando el productor: %v", err)
	}

	return &KafkaProducer{Producer: producer, Topic: topic}, nil
}

// Publicar un mensaje en Kafka
func (kp *KafkaProducer) PublishMessage(message []byte) error {
	_, _, err := kp.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kp.Topic,
		Value: sarama.ByteEncoder(message),
	})
	if err != nil {
		log.Printf("Error publicando mensaje: %v", err)
		return err
	}
	log.Printf("Mensaje publicado en el topico: %s", kp.Topic)
	return nil
}
