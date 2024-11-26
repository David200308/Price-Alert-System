package subscriptions

import (
	"log"

	"github.com/David200308/go-api/Scheduler/initializers"
	"github.com/David200308/go-api/Scheduler/models"
)

func GetSubscriptions(mode, symbol string) ([]models.Alert, error) {
	var alerts []models.Alert

	if err := initializers.DB.Where("symbol = ? AND type = ? AND status = ? AND is_alert = ?", symbol, mode, "active", false).
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

func GetSubscribedSymbols(mode string, c chan []string) {
	var symbols []string

	if err := initializers.DB.Model(&models.Subscribed{}).
		Select("symbol").
		Where("type = ?", mode).
		Group("symbol").
		Order("COUNT(*) DESC").
		Pluck("symbol", &symbols).Error; err != nil {
		log.Println("Error getting subscribed symbols:", err)
		return
	}

	c <- symbols
}
