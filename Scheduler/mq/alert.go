package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/David200308/go-api/Scheduler/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	StockPriceAlertQueueName  = "notification:stock:price:alert"
	CryptoPriceAlertQueueName = "notification:crypto:price:alert"
)

func PriceAlert(userUUID, mode, stockSymbol, direction string, price float64) error {
	messageText := fmt.Sprintf("%s Alert: %s going %s to %.2f", mode, stockSymbol, direction, price)

	notification := models.AlertNotification{
		UserUUID:  userUUID,
		Message:   messageText,
		Timestamp: time.Now().Unix(),
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal user notification: %w", err)
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	var queueName string
	switch mode {
	case "stock":
		queueName = StockPriceAlertQueueName
	case "crypto":
		queueName = CryptoPriceAlertQueueName
	default:
		return fmt.Errorf("invalid mode: %s", mode)
	}

	err = MQPublish(queueName, message)
	if err != nil {
		return fmt.Errorf("failed to publish alert to queue (%s): %w", queueName, err)
	}

	log.Printf("Price alert published: %s", messageText)
	return nil
}
