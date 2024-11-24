package models

import "gorm.io/gorm"

type AlertNotification struct {
	UserUUID  string `json:"user_uuid"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type Alert struct {
	gorm.Model
	AlertUUID string  `gorm:"not null;unique"`
	UserUUID  string  `gorm:"not null"`
	Type      string  `gorm:"not null"` // "stock", "crypto"
	Symbol    string  `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Operator  string  `gorm:"not null"` // ">", "<"
	Frequency string  `gorm:"not null"` // "once", "daily", "always"
	Status    string  `gorm:"not null"` // "active", "inactive",
	IsAlert   bool    `gorm:"not null"`
}
