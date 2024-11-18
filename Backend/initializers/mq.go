package initializers

import (
	"os"

	"github.com/streadway/amqp"
)

type MQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var MQInstance *MQ

func InitMQ() (*MQ, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func (mq *MQ) Close() {
	mq.Channel.Close()
	mq.Connection.Close()
}

func (mq *MQ) Publish(exchange, key string, body []byte) error {
	return mq.Channel.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (mq *MQ) Consume(queue, exchange, key string) (<-chan amqp.Delivery, error) {
	if err := mq.Channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		return nil, err
	}

	if _, err := mq.Channel.QueueDeclare(queue, true, false, false, false, nil); err != nil {
		return nil, err
	}

	if err := mq.Channel.QueueBind(queue, key, exchange, false, nil); err != nil {
		return nil, err
	}

	msgs, err := mq.Channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
