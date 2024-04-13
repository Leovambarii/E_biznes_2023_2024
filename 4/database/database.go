package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"myapp/models"
)

func Init() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("store.db"))
	if err != nil {
		panic("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{})
	if err != nil {
		panic("Failed to migrate database")
	}

	DB.Model(&models.Product{}).Association("Category")

	return DB
}

func LoadCategoryData(db *gorm.DB, data []models.Category) error {
	for _, category := range data {
		var existingCategory models.Category
		res := db.Where("name = ?", category.Name).First(&existingCategory)
		if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}

		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			if err := db.Create(&category).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func LoadProductData(db *gorm.DB, data []models.Product) error {
	for _, product := range data {
		var existingProduct models.Product
		res := db.Where("name = ? AND description = ? AND category_id = ?", product.Name, product.Description, product.CategoryID).First(&existingProduct)
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
