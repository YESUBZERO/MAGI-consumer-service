package kafka

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/YESUBZERO/consumer-service/internal/config"
	"github.com/YESUBZERO/consumer-service/internal/models"
	"github.com/YESUBZERO/consumer-service/internal/repository"
)

// WorkerPool define la cantidad de workers concurrentes
const WorkerPool = 5

type Consumer struct {
	cfg        *config.Config
	producer   *Producer
	repository *repository.ShipRepository
}

func NewConsumer(cfg *config.Config, producer *Producer, repo *repository.ShipRepository) *Consumer {
	return &Consumer{cfg: cfg, producer: producer, repository: repo}
}

func (c *Consumer) ConsumeMessages() {
	log.Println("Iniciando consumidor de Kafka...")

	consumer, err := sarama.NewConsumer(c.cfg.Kafka.Brokers, nil)
	if err != nil {
		log.Fatalf("Error creando consumidor de Kafka: %v", err)
	}

	topics := []string{c.cfg.Kafka.StaticTopic, c.cfg.Kafka.EnrichedTopic}
	messageChannel := make(chan *sarama.ConsumerMessage, 100)
	var wg sync.WaitGroup

	for i := 0; i < WorkerPool; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range messageChannel {
				var aisMsg models.Ship
				json.Unmarshal(msg.Value, &aisMsg)

				if msg.Topic == c.cfg.Kafka.StaticTopic && aisMsg.IMO != 0 {
					if !c.repository.ShipExists(aisMsg.IMO) {
						data, _ := json.Marshal(aisMsg)
						c.producer.SendMessage(c.cfg.Kafka.ScrapeTopic, string(data))
					}
				}

				if msg.Topic == c.cfg.Kafka.EnrichedTopic {
					// Guardar o actualizar los datos en la base de datos
					if err := c.repository.SaveShip(aisMsg); err != nil {
						log.Printf("Error guardando datos del barco: %v", err)
					}
				}
			}
		}()
	}

	for _, topic := range topics {
		pc, _ := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		for msg := range pc.Messages() {
			messageChannel <- msg
		}
	}

	close(messageChannel)
	wg.Wait()
}
