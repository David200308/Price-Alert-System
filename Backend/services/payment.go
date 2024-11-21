package services

import (
	"fmt"
	"log"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
	"github.com/David200308/go-api/Backend/services/payment"
)

func CreatePayment(paymentData *models.Payment) (string, error) {
	if err := initializers.DB.Create(&paymentData).Error; err != nil {
		log.Println("Error inserting payment into database:", err)
		return "", err
	}

	amount := paymentData.Amount
	currency := paymentData.Currency
	paymentUUId := paymentData.PaymentUUID

	switch paymentData.Method {
	case "stripe":
		session, err := payment.CreateStripePaymentSession(amount, currency, paymentUUId)
		if err != nil {
			fmt.Printf("Error creating session: %v\n", err)
			return "", err
		}
		fmt.Printf("Checkout session created: %s\n", session.URL)
		return session.URL, nil
	default:
		return "", fmt.Errorf("invalid payment method")
	}
}

func GetPaymentByPaymentUUID(paymentUUID string, userUUID string) (*models.Payment, error) {
	var payment models.Payment
	if err := initializers.DB.
		Where("payment_uuid = ? AND user_uuid = ?", paymentUUID, userUUID).
		First(&payment).
		Error; err != nil {
		log.Println("Error getting payment:", err)
		return nil, err
	}

	return &payment, nil
}

func GetPaymentByOrderUUID(orderUUID string, userUUID string) (*models.Payment, error) {
	var payment models.Payment
	if err := initializers.DB.
		Where("order_uuid = ? AND user_uuid = ?", orderUUID, userUUID).
		First(&payment).
		Error; err != nil {
		log.Println("Error getting payment:", err)
		return nil, err
	}

	return &payment, nil
}

func GetAllPaymentByUserUUID(userUUID string) ([]models.Payment, error) {
	var payments []models.Payment
	if err := initializers.DB.Where("user_uuid = ?", userUUID).
		Order("created_at DESC").
		Find(&payments).
		Error; err != nil {
		log.Println("Error getting payments:", err)
		return nil, err
	}

	return payments, nil
}

func UpdatePaymentStatus(userUUID, paymentUUID, status string, referenceId string) error {
	if status != "pending" && status != "success" && status != "failed" && status != "cancelled" {
		return fmt.Errorf("invalid status")
	}

	if err := initializers.DB.Model(&models.Payment{}).
		Where("payment_uuid = ? AND user_uuid = ?", paymentUUID, userUUID).
		Update("status", status).
		Error; err != nil {
		log.Println("Error updating payment status:", err)
		return err
	}

	return nil
}
