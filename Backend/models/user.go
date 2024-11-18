package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserUUID string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique;encrypted"`
	Username string `gorm:"not null;unique;encrypted"`
	Password string `gorm:"not null;hashed"`
	Status   string `gorm:"not null"` // "active", "inactive", "blocked", "pending"
}

type UserNotification struct {
	UserEmail string `json:"user_email"`
	UserUUID  string `json:"user_uuid"`
	Status    string `json:"status"`
}
