package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myapp/models"
	"myapp/scopes"
	"net/http"
	"strconv"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) GetCategories(c echo.Context) error {
	var categories []models.Category
	if err := cc.DB.Scopes(scopes.WithProducts()).Find(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, &categories)
}

func (cc *CategoryController) GetCategory(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid category ID"})
	}

	var category models.Category
	if err := cc.DB.Scopes(scopes.WithProducts(), scopes.ByID(id)).First(&category).Error; err != nil {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid category ID"})
	}

	var category models.Category
	if err := cc.DB.Scopes(scopes.WithProducts(), scopes.ByID(id)).First(&category).Error; err != nil {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid category ID"})
	}

	var category models.Category
	if err := cc.DB.Scopes(scopes.WithProducts(), scopes.ByID(id)).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "category not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	var products int64
	cc.DB.Model(&models.Product{}).Scopes(scopes.ByCategoryID(uint64(category.ID))).Count(&products)

	if products > 0 {
		return c.JSON(http.StatusConflict, "Cannot delete category with products")
	}

	cc.DB.Delete(&models.Category{}, id)

	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted"})
}
