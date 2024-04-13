package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myapp/models"
	"net/http"
)

type CartController struct {
	DB *gorm.DB
}

type CartRequest struct {
	ProductIDs []uint `json:"productIDs"`
}

func NewCartController(db *gorm.DB) *CartController {
	return &CartController{DB: db}
}

func (cc *CartController) GetCarts(c echo.Context) error {
	var carts []models.Cart
	if err := cc.DB.Preload("Products").Find(&carts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, carts)
}

func (cc *CartController) AddCart(c echo.Context) error {
	var cart models.Cart

	var req CartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	var totalCost float32
	for _, productID := range req.ProductIDs {
		var existingProduct models.Product
		if err := cc.DB.First(&existingProduct, productID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "product does not exist"})
		}

		totalCost += existingProduct.Price
		cart.Products = append(cart.Products, &existingProduct)
	}

	cart.TotalCost = totalCost

	if err := cc.DB.Create(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, cart)
}

func (cc *CartController) GetCart(c echo.Context) error {
	id := c.Param("id")

	var cart models.Cart
	if err := cc.DB.Preload("Products").First(&cart, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "cart not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, cart)
}

func (cc *CartController) EditCart(c echo.Context) error {
	id := c.Param("id")

	var cart models.Cart
	if err := cc.DB.Preload("Products").First(&cart, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "cart not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	var req CartRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad request"})
	}

	if err := cc.DB.Model(&cart).Association("Products").Clear(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	var totalCost float32
	for _, productID := range req.ProductIDs {
		var existingProduct models.Product
		if err := cc.DB.First(&existingProduct, productID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "product does not exist"})
		}

		if err := cc.DB.Model(&cart).Association("Products").Append(&existingProduct); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}

		totalCost += existingProduct.Price
	}

	cart.TotalCost = totalCost

	if err := cc.DB.Save(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	if err := cc.DB.Preload("Products").First(&cart, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, cart)
}

func (cc *CartController) DeleteCart(c echo.Context) error {
	id := c.Param("id")

	var cart models.Cart
	if err := cc.DB.First(&cart, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "cart not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	if err := cc.DB.Delete(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart deleted"})
}
