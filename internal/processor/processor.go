package processor

import (
	"encoding/json"
	"log"

	"github.com/YESUBZERO/consumer-service/internal/kafka"
)

// Estructura de un mensaje AIS
type AISMessage struct {
	MsgType int `json:"msg_type"`
	IMO     int `json:"imo"`
	// Otros campos omitidos
}

// Procesar un mensaje mensaje Kafka
func ProcessMessage(producer *kafka.KafkaProducer, message []byte) error {
	// Decodificar el mensaje
	var aisMessage AISMessage
	if err := json.Unmarshal(message, &aisMessage); err != nil {
		return err
	}

	// Filtrar mensajes relevantes
	if aisMessage.MsgType == 5 || aisMessage.MsgType == 24 {
		if aisMessage.IMO != 0 {
			log.Printf("Mensaje procesado: %+v", aisMessage)

			// Serializar el mensaje procesado
			ProcessedMessage, err := json.Marshal(aisMessage)
			if err != nil {
				return err
			}

			// Publicar mensaje en otro topico
			return producer.PublishMessage(ProcessedMessage)
		} else {
			log.Printf("Mensaje descartado: IMO no definido")
		}
	} else {
		log.Printf("Mensaje descartado: Tipo de mensaje: %d", aisMessage.MsgType)
	}

	return nil
}
