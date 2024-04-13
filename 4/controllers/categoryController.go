package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myapp/models"
	"net/http"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := cc.DB.Preload("Products").Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, &categories)
}

func (cc *CategoryController) GetCategory(c echo.Context) error {
	id := c.Param("id")

	var category models.Category
	if err := cc.DB.Preload("Products").First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, &category)
}

func (cc *CategoryController) AddCategory(c echo.Context) error {
	var category models.Category

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	if err := cc.DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) EditCategory(c echo.Context) error {
	id := c.Param("id")

	var category models.Category
	if err := cc.DB.Preload("Products").First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	var newCategory models.Category
	if err := c.Bind(&newCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	category.Name = newCategory.Name

	if err := cc.DB.Save(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	var products int64
	cc.DB.Model(&models.Product{}).Where("category_id = ?", category.ID).Count(&products)

	if products > 0 {
		return c.JSON(http.StatusConflict, "Cannot delete category with products")
	}

	cc.DB.Delete(&models.Category{}, id)

	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted"})
}
