package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    int32   `json:"category"`
	Price       float32 `json:"price"`
	Stock       int32   `json:"stock"`
}
