package processor

import (
	"encoding/json"
	"log"

	"github.com/YESUBZERO/consumer-service/internal/kafka"
)

// Estructura del mensaje AIS
type AISMessage struct {
	MsgType int `json:"msg_type"`
	IMO     int `json:"imo"`
}

// Procesar un mensaje Kafka
func ProcessMessage(producer *kafka.KafkaProducer, message []byte) error {
	var aisMessage AISMessage

	// Deserializar el mensaje
	if err := json.Unmarshal(message, &aisMessage); err != nil {
		return err
	}

	// Filtrar mensajes relevantes
	if aisMessage.MsgType == 5 || aisMessage.MsgType == 24 {
		if aisMessage.IMO != 0 {
			log.Printf("Procesando mensaje con IMO: %d", aisMessage.IMO)

			// Publicar mensaje procesado
			processedMessage, err := json.Marshal(aisMessage)
			if err != nil {
				return err
			}
			return producer.PublishMessage(processedMessage)
		}
		log.Println("Mensaje descartado: IMO inválido")
	} else {
		log.Printf("Mensaje descartado: Tipo %d", aisMessage.MsgType)
	}

	return nil
}
