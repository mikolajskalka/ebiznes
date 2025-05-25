package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikolajskalka/ebiznes/exercise4/controllers"
)

// Route path constants
const (
	ProductByIDPath  = "/products/:id"
	CategoryByIDPath = "/categories/:id"
)

// Configure all routes for the application
func SetupRoutes(e *echo.Echo) {
	// Product routes
	e.GET("/products", controllers.GetAllProducts)
	e.GET(ProductByIDPath, controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT(ProductByIDPath, controllers.UpdateProduct)
	e.DELETE(ProductByIDPath, controllers.DeleteProduct)
	e.GET("/products/category/:categoryId", controllers.GetProductsByCategory)
	e.GET("/products/price-range", controllers.GetProductsByPriceRange)

	// Category routes
	e.GET("/categories", controllers.GetAllCategories)
	e.GET("/categories/with-products", controllers.GetCategoriesWithProducts)
	e.GET(CategoryByIDPath, controllers.GetCategory)
	e.POST("/categories", controllers.CreateCategory)
	e.PUT(CategoryByIDPath, controllers.UpdateCategory)
	e.DELETE(CategoryByIDPath, controllers.DeleteCategory)
	e.GET("/categories/search", controllers.SearchCategoriesByName)

	// Cart routes
	e.GET("/carts", controllers.GetAllCarts)
	e.GET("/carts/:id", controllers.GetCart)
	e.POST("/carts", controllers.CreateCart)
	e.POST("/carts/:id/items", controllers.AddItemToCart)
	e.DELETE("/carts/:id/items/:itemId", controllers.RemoveItemFromCart)
	e.GET("/carts/user/:userId", controllers.GetCartByUser)
}
