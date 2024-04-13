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
	categoryData := []models.Category{
		{Name: "Furniture"},
		{Name: "Electronics"},
	}

	err := database.LoadCategoryData(DB, categoryData)
	if err != nil {
		panic(err)
	}

	productData := []models.Product{
		{Name: "Table", Description: "Made of wood.", CategoryID: 1, Price: 350.99, Stock: 15},
		{Name: "Table Lamp", Description: "Modern product.", CategoryID: 1, Price: 55.99, Stock: 30},
		{Name: "Laptop", Description: "Newest technology.", CategoryID: 2, Price: 1250.0, Stock: 5},
	}
	err = database.LoadProductData(DB, productData)
	if err != nil {
		panic(err)
	}

	productController := controllers.NewProductController(DB)
	e.GET("/products", productController.GetProducts)
	e.POST("/products", productController.AddProduct)
	e.GET("/products/:id", productController.GetProduct)
	e.PUT("/products/:id", productController.EditProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)

	cartController := controllers.NewCartController(DB)
	e.GET("/carts", cartController.GetCarts)
	e.POST("/carts", cartController.AddCart)
	e.GET("/carts/:id", cartController.GetCart)
	e.PUT("/carts/:id", cartController.EditCart)
	e.DELETE("/carts/:id", cartController.DeleteCart)

	categoryController := controllers.NewCategoryController(DB)

	e.GET("/categories", categoryController.GetCategories)
	e.POST("/categories", categoryController.AddCategory)
	e.GET("/categories/:id", categoryController.GetCategory)
	e.PUT("/categories/:id", categoryController.EditCategory)
	e.DELETE("/categories/:id", categoryController.DeleteCategory)

	e.Logger.Fatal(e.Start(":8080"))
}
