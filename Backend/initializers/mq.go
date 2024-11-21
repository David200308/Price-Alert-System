package initializers

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var MQInstance *MQ

func MQPublish(queueName string, message amqp.Publishing) error {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
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

func (mq *MQ) Close() {
	if mq.Channel != nil {
		mq.Channel.Close()
	}
	if mq.Connection != nil {
		mq.Connection.Close()
	}
	fmt.Println("RabbitMQ connection and channel closed")
}
