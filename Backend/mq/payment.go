package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

const PaymentCreatedQueueName = "notification:payment:created"
const PaymentSuccessQueueName = "notification:payment:success"
const PaymentCancelledQueueName = "notification:payment:cancelled"

func PaymentCreated(userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "created",
		Timestamp:   time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err = initializers.MQPublish(PaymentCreatedQueueName, message)
	if err != nil {
		return fmt.Errorf("failed to initialize MQ instance: %w", err)
	}

	log.Println("Payment created event published")

	return nil
}

func PaymentSuccessful(userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "success",
		Timestamp:   time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err = initializers.MQPublish(PaymentSuccessQueueName, message)
	if err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}

	return nil
}

func PaymentCancelled(userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "cancelled",
		Timestamp:   time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err = initializers.MQPublish(PaymentCancelledQueueName, message)
	if err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}

	return nil
}
