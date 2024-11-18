package mq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/David200308/go-api/Backend/models"
)

func UserCreated(ctx context.Context, userUUID string, email string) error {
	notification := models.UserNotification{
		UserEmail: email,
		UserUUID:  userUUID,
		Status:    "created",
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	if err := initializers.MQInstance.Publish("users.event.created", "User.Created", body); err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}
	return nil
}

func UserVerify(ctx context.Context, userUUID string, email string) error {
	notification := models.UserNotification{
		UserEmail: email,
		UserUUID:  userUUID,
		Status:    "verify",
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal payment success message: %w", err)
	}

	if err := initializers.MQInstance.Publish("users.event.verify", "User.Verify", body); err != nil {
		return fmt.Errorf("failed to publish payment success message: %w", err)
	}
	return nil
}
