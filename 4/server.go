package main

import (
	"myapp/controllers"
	"myapp/database"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	DB := database.Init()
	productData := []models.Product{
		{Name: "Table", Description: "Made of wood.", Category: 1, Price: 350.99, Stock: 15},
		{Name: "Table Lamp", Description: "Modern product.", Category: 1, Price: 55.99, Stock: 30},
		{Name: "Laptop", Description: "Newest technology.", Category: 2, Price: 1250.0, Stock: 5},
	}
	err := database.LoadData(DB, productData)
	if err != nil {
		panic(err)
	}

	productController := controllers.NewProductController(DB)
	e.GET("/products", productController.GetProducts)
	e.POST("/products", productController.AddProduct)
	e.GET("/products/:id", productController.GetProduct)
	e.PUT("/products/:id", productController.EditProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
