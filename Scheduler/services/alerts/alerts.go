package alerts

import (
	"log"

	"github.com/David200308/go-api/Scheduler/mq"
	"github.com/David200308/go-api/Scheduler/services/prices"
	"github.com/David200308/go-api/Scheduler/services/subscriptions"
)

func SendPriceAlertNotification(mode, symbol string) {
	var price float64
	var err error

	switch mode {
	case "stock":
		price, err = prices.GetStockPrice(symbol)
		if err != nil {
			log.Printf("Error fetching %s price for %s: %v", mode, symbol, err)
			return
		}
	case "crypto":
		price, err = prices.GetCryptoPrice(symbol)
		if err != nil {
			log.Printf("Error fetching %s price for %s: %v", mode, symbol, err)
			return
		}
	}

	alerts, err := subscriptions.GetSubscriptions(mode, symbol)
	if err != nil {
		log.Printf("Error fetching subscriptions for %s %s: %v", mode, symbol, err)
		return
	}

	for _, alert := range alerts {
		var shouldNotify bool

		switch alert.Operator {
		case ">":
			shouldNotify = price > alert.Price
		case "<":
			shouldNotify = price < alert.Price
		default:
			log.Printf("Unknown operator %s in alert for %s", alert.Operator, symbol)
			continue
		}

		if !shouldNotify {
			continue
		}

		direction := "up"
		if alert.Operator == "<" {
			direction = "down"
		}
		err = mq.PriceAlert(alert.UserUUID, mode, symbol, direction, price)
		if err != nil {
			log.Printf("Failed to send alert notification for %s: %v", symbol, err)
			continue
		}

		log.Printf("Alert notification sent for %s to user %s", symbol, alert.UserUUID)

		if alert.Frequency == "once" {
			alert.IsAlert = true
			alert.Status = "inactive"

			if err := subscriptions.UpdateAlertStatus(&alert); err != nil {
				log.Printf("Failed to update alert status for %s: %v", symbol, err)
			}
		}

		if alert.Frequency == "daily" {
			alert.IsAlert = true

			if err := subscriptions.UpdateAlertStatus(&alert); err != nil {
				log.Printf("Failed to update alert status for %s: %v", symbol, err)
			}
		}
	}
}
