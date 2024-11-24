package main

import (
	"log"

	"github.com/David200308/go-api/Scheduler/initializers"
	"github.com/David200308/go-api/Scheduler/models"
)

func UpdateAlertStatus() error {
	if err := initializers.DB.Model(&models.Alert{}).
		Where("status = ? AND frequency = ? AND is_alert = ?", "active", "daily", "true").
		Update("is_alert", "false").
		Error; err != nil {
		log.Println("Error updating alert status:", err)
		return err
	}
	return nil
}

func main() {
	if err := UpdateAlertStatus(); err != nil {
		log.Println("Error updating alert status:", err)
	}
}
