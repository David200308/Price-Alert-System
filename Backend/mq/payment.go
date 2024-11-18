package mq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
)

func PaymentCreated(ctx context.Context, userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "created",
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	if err := initializers.MQInstance.Publish("payments.event.created", "Payment.Created", body); err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}
	return nil
}

func PaymentSuccessful(ctx context.Context, userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "success",
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	if err := initializers.MQInstance.Publish("payments.event.success", "Payment.Success", body); err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}
	return nil
}

func PaymentCancelled(ctx context.Context, userUUID string, paymentUUID string) error {
	notification := models.PaymentNotification{
		PaymentUUID: paymentUUID,
		UserUUID:    userUUID,
		Status:      "cancelled",
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	if err := initializers.MQInstance.Publish("payments.event.cancelled", "Payment.Cancelled", body); err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}
	return nil
}
