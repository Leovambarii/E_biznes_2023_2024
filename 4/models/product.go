package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  uint64  `json:"category_id"`
	Price       float64 `json:"price"`
	Stock       uint64  `json:"stock"`
}
