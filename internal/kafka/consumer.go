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
	log.Println("🚀 Iniciando consumidor de Kafka...")

	consumer, err := sarama.NewConsumer(c.cfg.Kafka.Brokers, nil)
	if err != nil {
		log.Fatalf("❌ Error creando consumidor de Kafka: %v", err)
	}
	defer consumer.Close()

	topics := []string{c.cfg.Kafka.StaticTopic, c.cfg.Kafka.EnrichedTopic}
	messageChannel := make(chan *sarama.ConsumerMessage, 100)
	var wg sync.WaitGroup

	// 🔄 Iniciar pool de workers
	for i := 0; i < WorkerPool; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for msg := range messageChannel {
				var aisMsg models.Ship

				json.Unmarshal(msg.Value, &aisMsg)

				// ✅ Procesar mensaje de static-message
				if msg.Topic == c.cfg.Kafka.StaticTopic && aisMsg.IMO != 0 {
					if !c.repository.ShipExists(aisMsg.IMO) {
						//log.Printf("📡 [Worker %d] Enviando barco con IMO %d a Scraper...", workerID, aisMsg.IMO)
						data, _ := json.Marshal(aisMsg)
						c.producer.SendMessage(c.cfg.Kafka.ScrapeTopic, string(data))
					}
				}

				// ✅ Procesar mensaje de enriched-message y guardar en BD
				if msg.Topic == c.cfg.Kafka.EnrichedTopic {
					log.Printf("📦 [Worker %d] Guardando barco con IMO %d.", workerID, aisMsg.IMO)
					if err := c.repository.SaveShip(aisMsg); err != nil {
						log.Printf("❌ [Worker %d] Error guardando en BD: %v\nDatos: %+v", workerID, err, aisMsg)
					}
				}
			}
		}(i)
	}

	// Consumir mensajes de Kafka y enviarlos a messageChannel
	for _, topic := range topics {
		log.Printf("🔗 Suscribiéndose al tópico: %s", topic)

		partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("❌ Error consumiendo el tópico %s: %v", topic, err)
		}

		go func(topic string, pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//log.Printf("📨 Mensaje recibido en %s: %s", topic, string(msg.Value))
				messageChannel <- msg
			}
		}(topic, partitionConsumer)
	}

	// NO cierres messageChannel aquí, deja que los workers lo manejen
	wg.Wait()

}
