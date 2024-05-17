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
	if err := cc.DB.Scopes(scopes.WithProducts()).Find(&carts).Error; err != nil {
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

	var totalCost float64
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid cart ID"})
	}

	var cart models.Cart
	if err := cc.DB.Scopes(scopes.WithProducts(), scopes.ByID(id)).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "cart not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, cart)
}

func (cc *CartController) EditCart(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid cart ID"})
	}

	var cart models.Cart
	if err := cc.DB.Scopes(scopes.WithProducts(), scopes.ByID(id)).First(&cart).Error; err != nil {
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

	var totalCost float64
	for _, productID := range req.ProductIDs {
		var existingProduct models.Product
		if err := cc.DB.Scopes(scopes.ByID(uint64(productID))).First(&existingProduct, productID).Error; err != nil {
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid cart ID"})
	}

	var cart models.Cart
	if err := cc.DB.Scopes(scopes.ByID(id)).First(&cart).Error; err != nil {
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
