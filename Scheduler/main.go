package main

import (
	"github.com/David200308/go-api/Scheduler/initializers"
	"github.com/David200308/go-api/Scheduler/services/alerts"
	"github.com/David200308/go-api/Scheduler/services/subscriptions"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

	subscribedStockChannel := make(chan []string)
	subscribedCryptoChannel := make(chan []string)

	go subscriptions.GetSubscribedSymbols("stock", subscribedStockChannel)
	go subscriptions.GetSubscribedSymbols("crypto", subscribedCryptoChannel)

	subscribedStock, subscribedCrypto := <-subscribedStockChannel, <-subscribedCryptoChannel

	go func() {
		for _, symbol := range subscribedStock {
			go alerts.SendPriceAlertNotification("stock", symbol)
		}
	}()

	go func() {
		for _, symbol := range subscribedCrypto {
			go alerts.SendPriceAlertNotification("crypto", symbol)
		}
	}()

}
