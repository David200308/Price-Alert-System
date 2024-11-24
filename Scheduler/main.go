package main

import (
	"github.com/David200308/go-api/Scheduler/initializers"
	"github.com/David200308/go-api/Scheduler/services/alerts"
	"github.com/David200308/go-api/Scheduler/services/subscriptions"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

	stocks, err := subscriptions.GetSubscribedStocks()
	if err != nil {
		return
	}

	for _, stock := range stocks {
		go alerts.SendStockAlertNotification(stock)
	}
}
