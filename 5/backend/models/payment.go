package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CardNumber string  `json:"cardNumber"`
	ExpiryDate string  `json:"expiryDate"`
	Amount     float64 `json:"amount"`
}
