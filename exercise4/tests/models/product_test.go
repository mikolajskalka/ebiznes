package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

func TestProductModel(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test products
	db.Unscoped().Where("name LIKE ?", "Test Product%").Delete(&models.Product{})

	// Test 1: Create a new product
	t.Run("Create Product", func(t *testing.T) {
		product := models.Product{
			Name:        "Test Product 1",
			Description: "Test Description 1",
			Price:       99.99,
			Quantity:    10,
			CategoryID:  1,
		}

		result := db.Create(&product)
		assert.Nil(t, result.Error)
		assert.NotZero(t, product.ID)
		assert.Equal(t, "Test Product 1", product.Name)
		assert.Equal(t, 99.99, product.Price)
	})

	// Test 2: Find a product
	t.Run("Find Product", func(t *testing.T) {
		var product models.Product
		result := db.Where("name = ?", "Test Product 1").First(&product)
		assert.Nil(t, result.Error)
		assert.Equal(t, "Test Product 1", product.Name)
		assert.Equal(t, "Test Description 1", product.Description)
		assert.Equal(t, 99.99, product.Price)
	})

	// Test 3: Update a product
	t.Run("Update Product", func(t *testing.T) {
		var product models.Product
		db.Where("name = ?", "Test Product 1").First(&product)

		product.Price = 129.99
		product.Description = "Updated Description"

		result := db.Save(&product)
		assert.Nil(t, result.Error)

		var updatedProduct models.Product
		db.Where("name = ?", "Test Product 1").First(&updatedProduct)
		assert.Equal(t, 129.99, updatedProduct.Price)
		assert.Equal(t, "Updated Description", updatedProduct.Description)
	})

	// Test 4: Delete a product (soft delete)
	t.Run("Delete Product", func(t *testing.T) {
		var product models.Product
		db.Where("name = ?", "Test Product 1").First(&product)

		result := db.Delete(&product)
		assert.Nil(t, result.Error)

		var deletedProduct models.Product
		result = db.Where("name = ?", "Test Product 1").First(&deletedProduct)
		assert.Error(t, result.Error) // Should not find the deleted product

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().Where("name = ?", "Test Product 1").First(&deletedProduct)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedProduct.DeletedAt.Time)
	})

	// Test 5: Test ActiveProducts scope
	t.Run("ActiveProducts Scope", func(t *testing.T) {
		// Create two products
		product1 := models.Product{
			Name:        "Test Product Active",
			Description: "Active product",
			Price:       19.99,
			Quantity:    5,
			CategoryID:  1,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Test Product To Delete",
			Description: "Will be deleted",
			Price:       29.99,
			Quantity:    3,
			CategoryID:  1,
		}
		db.Create(&product2)

		// Delete one product
		db.Delete(&product2)

		// Test ActiveProducts scope
		var products []models.Product
		db.Scopes(models.ActiveProducts).Where("name LIKE ?", "Test Product%").Find(&products)

		// Should only find non-deleted products
		activeFound := false
		deletedFound := false

		for _, p := range products {
			if p.Name == "Test Product Active" {
				activeFound = true
			}
			if p.Name == "Test Product To Delete" {
				deletedFound = true
			}
		}

		assert.True(t, activeFound)
		assert.False(t, deletedFound)

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
	})

	// Test 6: Test InStock scope
	t.Run("InStock Scope", func(t *testing.T) {
		// Create products with different stock levels
		product1 := models.Product{
			Name:        "Test Product In Stock",
			Description: "Has stock",
			Price:       39.99,
			Quantity:    10,
			CategoryID:  1,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Test Product Out of Stock",
			Description: "No stock",
			Price:       49.99,
			Quantity:    0,
			CategoryID:  1,
		}
		db.Create(&product2)

		// Test InStock scope
		var products []models.Product
		db.Scopes(models.InStock).Where("name LIKE ?", "Test Product%").Find(&products)

		// Should only find products with stock
		inStockFound := false
		outOfStockFound := false

		for _, p := range products {
			if p.Name == "Test Product In Stock" {
				inStockFound = true
			}
			if p.Name == "Test Product Out of Stock" {
				outOfStockFound = true
			}
		}

		assert.True(t, inStockFound)
		assert.False(t, outOfStockFound)

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
	})

	// Test 7: Test ByCategoryID scope
	t.Run("ByCategoryID Scope", func(t *testing.T) {
		// Create products with different categories
		product1 := models.Product{
			Name:        "Test Product Cat 1",
			Description: "Category 1",
			Price:       59.99,
			Quantity:    8,
			CategoryID:  1,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Test Product Cat 2",
			Description: "Category 2",
			Price:       69.99,
			Quantity:    6,
			CategoryID:  2,
		}
		db.Create(&product2)

		// Test ByCategoryID scope
		var products []models.Product
		db.Scopes(models.ByCategoryID(1)).Where("name LIKE ?", "Test Product Cat%").Find(&products)

		// Should only find products with CategoryID = 1
		cat1Found := false
		cat2Found := false

		for _, p := range products {
			if p.Name == "Test Product Cat 1" {
				cat1Found = true
			}
			if p.Name == "Test Product Cat 2" {
				cat2Found = true
			}
		}

		assert.True(t, cat1Found)
		assert.False(t, cat2Found)

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
	})

	// Test 8: Test ByPriceRange scope
	t.Run("ByPriceRange Scope", func(t *testing.T) {
		// Create products with different price ranges
		product1 := models.Product{
			Name:        "Test Product Low Price",
			Description: "Low price range",
			Price:       10.99,
			Quantity:    5,
			CategoryID:  1,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Test Product Mid Price",
			Description: "Medium price range",
			Price:       50.99,
			Quantity:    5,
			CategoryID:  1,
		}
		db.Create(&product2)

		product3 := models.Product{
			Name:        "Test Product High Price",
			Description: "High price range",
			Price:       100.99,
			Quantity:    5,
			CategoryID:  1,
		}
		db.Create(&product3)

		// Test ByPriceRange scope (30-80)
		var products []models.Product
		db.Scopes(models.ByPriceRange(30.0, 80.0)).Where("name LIKE ?", "Test Product%").Find(&products)

		// Should only find products in the given price range
		lowFound := false
		midFound := false
		highFound := false

		for _, p := range products {
			if p.Name == "Test Product Low Price" {
				lowFound = true
			}
			if p.Name == "Test Product Mid Price" {
				midFound = true
			}
			if p.Name == "Test Product High Price" {
				highFound = true
			}
		}

		assert.False(t, lowFound)
		assert.True(t, midFound)
		assert.False(t, highFound)

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
		db.Unscoped().Delete(&product3)
	})
}
