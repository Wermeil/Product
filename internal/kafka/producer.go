package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

// Producer отправляет сообщения в Kafka
type Producer struct {
	writer *kafka.Writer
}

// NewProducer создает нового продюсера
func NewProducer(broker string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(broker), // Адрес Kafka сервера
			Topic: "user-events",     // Топик по умолчанию
		},
	}
}

// SendMessage отправляет одно сообщение
func (p *Producer) SendMessage(ctx context.Context, eventType string, data interface{}) error {
	// Преобразуем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Отправляем сообщение в Kafka
	return p.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(eventType), // Ключ = тип события (например "user.created")
			Value: jsonData,          // Значение = данные события в JSON
		},
	)
}
