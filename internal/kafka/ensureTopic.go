package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

// Crear tópicos en Kafka si no existen
func EnsureTopicExists(brokers []string, topic string, partitions, replicationFactor int) error {
	admin, err := sarama.NewClusterAdmin(brokers, nil)
	if err != nil {
		return err
	}
	defer admin.Close()

	// Verificar si el tópico ya existe
	topics, err := admin.ListTopics()
	if err != nil {
		return err
	}
	if _, exists := topics[topic]; exists {
		log.Printf("El tópico '%s' ya existe", topic)
		return nil
	}

	// Crear el tópico si no existe
	err = admin.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     int32(partitions),
		ReplicationFactor: int16(replicationFactor),
	}, false)
	if err != nil {
		return err
	}
	log.Printf("Tópico '%s' creado con éxito", topic)
	return nil
}
