package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// Consumer читает сообщения из Kafka
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer создает нового консюмера
func NewConsumer(broker, topic, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker}, // Адрес Kafka
			Topic:   topic,            // Какой топик слушать
			GroupID: groupID,          // Группа консюмеров (для балансировки)
		}),
	}
}

// StartListening начинает слушать сообщения
func (c *Consumer) StartListening(ctx context.Context) {
	for {
		// Читаем следующее сообщение
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Ошибка чтения: %v", err)
			continue
		}

		// Выводим что получили
		log.Printf("Получено сообщение: ключ=%s, значение=%s",
			string(msg.Key), string(msg.Value))
	}
}
