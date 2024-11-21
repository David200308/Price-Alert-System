package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PaymentUUID string `gorm:"not null;unique"`
	OrderUUID   string `gorm:"not null"`
	UserUUID    string `gorm:"not null"`
	Amount      int64  `gorm:"not null"`
	Currency    string `gorm:"not null"`
	Method      string `gorm:"not null"` // "stripe"
	Status      string `gorm:"not null"` // "pending", "success", "failed", "cancelled"
	ReferenceID string `gorm:"default:null"`
}

type PaymentNotification struct {
	PaymentUUID string `json:"payment_uuid"`
	UserUUID    string `json:"user_uuid"`
	Status      string `json:"status"`
	Timestamp   int64  `json:"timestamp"`
}
