package kafka_handler

import (
	"api-gateway/variables"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func SendToKafka(sendTopic string, sendKey string, sendValue string) {
	kafkaWriter := NewKafkaWriter(sendTopic)
	err := kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(sendKey),
			Value: []byte(sendValue),
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

func NewKafkaReader(topic string, group string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{variables.KAFKA_HOST},
		Topic:     topic,
		// Partition: 0,
		MinBytes:  1,    // 0
		MaxBytes:  10e6, // 10MB
	})
}
