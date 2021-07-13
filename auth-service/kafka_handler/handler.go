package kafka_handler

import (
	"auth-service/variables"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func SendToKafka(kafkaWriter kafka.Writer, sendKey []byte, sendValue []byte) {

	err := kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   sendKey,
			Value: sendValue,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := kafkaWriter.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
func NewKafkaWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(variables.KAFKA_HOST),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{variables.KAFKA_HOST},
		Topic:    topic,
		GroupID:  variables.KAFKA_GROUP,
		MinBytes: 0,    // 0
		MaxBytes: 10e6, // 10MB
	})
}
