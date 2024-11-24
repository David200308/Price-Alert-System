package mq

import (
	"fmt"
	"os"

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
