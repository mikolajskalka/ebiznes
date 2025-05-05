package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
)

// GetAllProducts - Get all products
func GetAllProducts(c echo.Context) error {
	var products []models.Product

	db := database.GetDB()
	result := db.Scopes(models.ActiveProducts).Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

// GetProduct - Get a product by ID
func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	var product models.Product
	db := database.GetDB()
	result := db.First(&product, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	return c.JSON(http.StatusOK, product)
}

// CreateProduct - Create a new product
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product data",
		})
	}

	db := database.GetDB()
	result := db.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, product)
}

// UpdateProduct - Update an existing product
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	// Check if product exists
	var existingProduct models.Product
	db := database.GetDB()
	result := db.First(&existingProduct, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	// Bind updated data
	updatedProduct := new(models.Product)
	if err := c.Bind(updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product data",
		})
	}

	// Update product
	existingProduct.Name = updatedProduct.Name
	existingProduct.Description = updatedProduct.Description
	existingProduct.Price = updatedProduct.Price
	existingProduct.Quantity = updatedProduct.Quantity
	existingProduct.CategoryID = updatedProduct.CategoryID

	db.Save(&existingProduct)

	return c.JSON(http.StatusOK, existingProduct)
}

// DeleteProduct - Delete a product
func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	// Check if product exists
	var product models.Product
	db := database.GetDB()
	result := db.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	// Delete product (soft delete with GORM)
	db.Delete(&product)

	return c.NoContent(http.StatusNoContent)
}

// GetProductsByCategory - Get products by category ID using scopes
func GetProductsByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	var products []models.Product
	db := database.GetDB()
	result := db.Scopes(models.ActiveProducts, models.ByCategoryID(uint(categoryID))).Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

// GetProductsByPriceRange - Get products by price range using scopes
func GetProductsByPriceRange(c echo.Context) error {
	minStr := c.QueryParam("min")
	maxStr := c.QueryParam("max")

	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		min = 0
	}

	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil {
		max = 1000000 // Default high value
	}

	var products []models.Product
	db := database.GetDB()
	result := db.Scopes(models.ActiveProducts, models.ByPriceRange(min, max)).Find(&products)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve products",
		})
	}

	return c.JSON(http.StatusOK, products)
}
