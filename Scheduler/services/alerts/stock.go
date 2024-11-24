package alerts

import (
	"log"

	"github.com/David200308/go-api/Scheduler/mq"
	"github.com/David200308/go-api/Scheduler/services/prices"
	"github.com/David200308/go-api/Scheduler/services/subscriptions"
)

func SendStockAlertNotification(stockSymbol string) {
	stockPrice, err := prices.GetStockPrice(stockSymbol)
	if err != nil {
		log.Printf("Error fetching stock price for %s: %v", stockSymbol, err)
		return
	}

	alerts, err := subscriptions.GetStockSubscriptions(stockSymbol)
	if err != nil {
		log.Printf("Error fetching subscriptions for stock %s: %v", stockSymbol, err)
		return
	}

	for _, alert := range alerts {
		var shouldNotify bool

		switch alert.Operator {
		case ">":
			shouldNotify = stockPrice > alert.Price
		case "<":
			shouldNotify = stockPrice < alert.Price
		default:
			log.Printf("Unknown operator %s in alert for %s", alert.Operator, stockSymbol)
			continue
		}

		if !shouldNotify {
			continue
		}

		direction := "up"
		if alert.Operator == "<" {
			direction = "down"
		}
		err = mq.PriceAlert(alert.UserUUID, "stock", stockSymbol, direction, stockPrice)
		if err != nil {
			log.Printf("Failed to send alert notification for %s: %v", stockSymbol, err)
			continue
		}

		log.Printf("Alert notification sent for %s to user %s", stockSymbol, alert.UserUUID)

		if alert.Frequency == "once" {
			alert.IsAlert = true
			alert.Status = "inactive"

			if err := subscriptions.UpdateAlertStatus(&alert); err != nil {
				log.Printf("Failed to update alert status for %s: %v", stockSymbol, err)
			}
		}

		if alert.Frequency == "daily" {
			alert.IsAlert = true

			if err := subscriptions.UpdateAlertStatus(&alert); err != nil {
				log.Printf("Failed to update alert status for %s: %v", stockSymbol, err)
			}
		}
	}
}
