package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myapp/models"
)

func Init() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("product.db"))
	if err != nil {
		panic("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return DB
}

func LoadData(db *gorm.DB, weatherData []models.Product) error {
	for _, product := range weatherData {
		var existingProduct models.Product
		res := db.Where("name = ? AND description = ? AND category = ?", product.Name, product.Description, product.Category).First(&existingProduct)
		if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&product).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
