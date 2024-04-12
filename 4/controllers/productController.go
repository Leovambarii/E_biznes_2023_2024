package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myapp/models"
	"net/http"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) GetProducts(c echo.Context) error {
	var products []models.Product
	if err := pc.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) AddProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		println("Error AddProduct Bind: " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	if product.Name == "" || product.Description == "" || product.Category == 0 || product.Price == 0 || product.Stock == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing required fields"})
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		println("Error AddProduct Create: " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) EditProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	newProduct := new(models.Product)
	if err := c.Bind(newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	if newProduct.Name != "" {
		product.Name = newProduct.Name
	}
	if newProduct.Description != "" {
		product.Description = newProduct.Description
	}
	if newProduct.Category != 0 {
		product.Category = newProduct.Category
	}
	if newProduct.Price != 0 {
		product.Price = newProduct.Price
	}
	if newProduct.Stock != 0 {
		product.Stock = newProduct.Stock
	}

	if err := pc.DB.Save(&product).Error; err != nil {
		println("Error editProduct Save: " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	if err := pc.DB.Delete(&product).Error; err != nil {
		println("Error DeleteProduct: " + err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted"})
}
