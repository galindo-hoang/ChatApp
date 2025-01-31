package services

import (
	"fmt"
	"github.com/ChatService/internal/configs"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitLoggerClient(kafkaConfig configs.Kafka) (*kafka.Producer, func(), error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", kafkaConfig.Host, kafkaConfig.Port),
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return nil, nil, err
	}

	cleanup := func() {
		fmt.Println("Clean-up database...")
		p.Close()
	}
	return p, cleanup, nil
}
