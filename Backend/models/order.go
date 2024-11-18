package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderUUID string `gorm:"not null;unique"`
	UserUUID  string `gorm:"not null"`
	Amount    int64  `gorm:"not null"`
	Currency  string `gorm:"not null"`
	Status    string `gorm:"not null"` // "unpaid", "paid", "cancelled"
}
