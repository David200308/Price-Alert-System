package models

import (
	"gorm.io/gorm"
)

type Subscribed struct {
	gorm.Model
	Type   string `gorm:"not null"` // "stock", "crypto"
	Symbol string `gorm:"not null"`
	Count  int    `gorm:"not null"`
}
