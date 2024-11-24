package initializers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MQPublish(queueName string, message amqp.Publishing) error {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		return fmt.Errorf("RABBITMQ_URL environment variable is not set")
	}

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	if err := channel.Publish("", queueName, false, false, message); err != nil {
		return fmt.Errorf("failed to publish user notification: %w", err)
	}

	return nil
}

func MQConsume(queueName string, timeout time.Duration) ([]string, error) {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		return nil, fmt.Errorf("RABBITMQ_URL environment variable is not set")
	}

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	defer channel.Close()

	messages, err := channel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var res []string

	for {
		select {
		case message, ok := <-messages:
			if !ok {
				log.Println("Message channel closed")
				return res, nil
			}
			log.Printf("Message: %s\n", message.Body)
			res = append(res, string(message.Body))

		case <-ctx.Done():
			log.Println("Timeout reached or context cancelled")
			return res, nil
		}
	}
}
