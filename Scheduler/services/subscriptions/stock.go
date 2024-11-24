package subscriptions

import (
	"log"

	"github.com/David200308/go-api/Scheduler/initializers"
	"github.com/David200308/go-api/Scheduler/models"
)

func GetStockSubscriptions(stockSymbol string) ([]models.Alert, error) {
	var alerts []models.Alert

	if err := initializers.DB.Where("symbol = ? AND type = ? AND status = ? AND is_alert = ?", stockSymbol, "stock", "active", false).
		Find(&alerts).Error; err != nil {
		log.Println("Error getting stock subscriptions:", err)
		return nil, err
	}

	return alerts, nil
}

func UpdateAlertStatus(alert *models.Alert) error {
	if err := initializers.DB.Save(alert).Error; err != nil {
		log.Println("Error updating alert status:", err)
		return err
	}

	return nil
}

func GetSubscribedStocks() ([]string, error) {
	var stockSymbols []string

	if err := initializers.DB.Model(&models.Subscribed{}).
		Distinct("symbol").
		Where("type = ?", "stock").
		Order("count DESC").
		Pluck("symbol", &stockSymbols).Error; err != nil {
		log.Println("Error getting subscribed stocks:", err)
		return nil, err
	}

	return stockSymbols, nil
}
