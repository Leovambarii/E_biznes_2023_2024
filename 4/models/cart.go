package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	TotalCost float64    `json:"total_cost"`
	Products  []*Product `gorm:"many2many:cart_products;" json:"products"`
}
