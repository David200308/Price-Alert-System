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

func UserCreated(userUUID string, email string) error {
	notification := models.UserNotification{
		UserEmail: email,
		UserUUID:  userUUID,
		Status:    "created",
		Timestamp: time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal user notification: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err = initializers.MQPublish(UserCreatedQueueName, message)
	if err != nil {
		return fmt.Errorf("failed to initialize MQ instance: %w", err)
	}
	log.Println("User created event published")

	return nil
}

func UserVerify(userUUID string, email string) error {
	notification := models.UserNotification{
		UserEmail: email,
		UserUUID:  userUUID,
		Status:    "verified",
		Timestamp: time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	}

	err = initializers.MQPublish(UserVerifiedQueueName, message)
	if err != nil {
		return fmt.Errorf("failed to initialize MQ instance: %w", err)
	}

	log.Println("User verified event published")

	return nil
}
