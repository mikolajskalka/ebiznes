package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

func TestCategoryModel(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test categories
	db.Unscoped().Where("name LIKE ?", "Test Category%").Delete(&models.Category{})

	// Test 1: Create a new category
	t.Run("Create Category", func(t *testing.T) {
		category := models.Category{
			Name:        "Test Category 1",
			Description: "Test Category Description 1",
		}

		result := db.Create(&category)
		assert.Nil(t, result.Error)
		assert.NotZero(t, category.ID)
		assert.Equal(t, "Test Category 1", category.Name)
		assert.Equal(t, "Test Category Description 1", category.Description)
	})

	// Test 2: Find a category
	t.Run("Find Category", func(t *testing.T) {
		var category models.Category
		result := db.Where("name = ?", "Test Category 1").First(&category)
		assert.Nil(t, result.Error)
		assert.Equal(t, "Test Category 1", category.Name)
		assert.Equal(t, "Test Category Description 1", category.Description)
	})

	// Test 3: Update a category
	t.Run("Update Category", func(t *testing.T) {
		var category models.Category
		db.Where("name = ?", "Test Category 1").First(&category)

		category.Name = "Updated Test Category"
		category.Description = "Updated Category Description"

		result := db.Save(&category)
		assert.Nil(t, result.Error)

		var updatedCategory models.Category
		db.First(&updatedCategory, category.ID)
		assert.Equal(t, "Updated Test Category", updatedCategory.Name)
		assert.Equal(t, "Updated Category Description", updatedCategory.Description)
	})

	// Test 4: Delete a category (soft delete)
	t.Run("Delete Category", func(t *testing.T) {
		var category models.Category
		db.Where("name = ?", "Updated Test Category").First(&category)

		result := db.Delete(&category)
		assert.Nil(t, result.Error)

		var deletedCategory models.Category
		result = db.Where("name = ?", "Updated Test Category").First(&deletedCategory)
		assert.Error(t, result.Error) // Should not find the deleted category

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().Where("name = ?", "Updated Test Category").First(&deletedCategory)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedCategory.DeletedAt.Time)
	})

	// Test 5: Test ActiveCategories scope
	t.Run("ActiveCategories Scope", func(t *testing.T) {
		// Create two categories
		category1 := models.Category{
			Name:        "Test Category Active",
			Description: "Active category",
		}
		db.Create(&category1)

		category2 := models.Category{
			Name:        "Test Category To Delete",
			Description: "Will be deleted",
		}
		db.Create(&category2)

		// Delete one category
		db.Delete(&category2)

		// Test ActiveCategories scope
		var categories []models.Category
		db.Scopes(models.ActiveCategories).Where("name LIKE ?", "Test Category%").Find(&categories)

		// Should only find non-deleted categories
		activeFound := false
		deletedFound := false

		for _, c := range categories {
			if c.Name == "Test Category Active" {
				activeFound = true
			}
			if c.Name == "Test Category To Delete" {
				deletedFound = true
			}
		}

		assert.True(t, activeFound)
		assert.False(t, deletedFound)

		// Clean up
		db.Unscoped().Delete(&category1)
		db.Unscoped().Delete(&category2)
	})

	// Test 6: Test ByName scope
	t.Run("ByName Scope", func(t *testing.T) {
		// Create categories with different names
		category1 := models.Category{
			Name:        "Test Category Apple",
			Description: "Apple products",
		}
		db.Create(&category1)

		category2 := models.Category{
			Name:        "Test Category Banana",
			Description: "Banana products",
		}
		db.Create(&category2)

		category3 := models.Category{
			Name:        "Test Category Apple Premium",
			Description: "Premium apple products",
		}
		db.Create(&category3)

		// Test ByName scope
		var categories []models.Category
		db.Scopes(models.ByName("Apple")).Where("name LIKE ?", "Test Category%").Find(&categories)

		// Should find categories containing 'Apple'
		appleFound := 0
		bananaFound := false

		for _, c := range categories {
			if c.Name == "Test Category Apple" || c.Name == "Test Category Apple Premium" {
				appleFound++
			}
			if c.Name == "Test Category Banana" {
				bananaFound = true
			}
		}

		assert.Equal(t, 2, appleFound)
		assert.False(t, bananaFound)

		// Clean up
		db.Unscoped().Delete(&category1)
		db.Unscoped().Delete(&category2)
		db.Unscoped().Delete(&category3)
	})

	// Test 7: Test WithProducts scope
	t.Run("WithProducts Scope", func(t *testing.T) {
		// Create a test category
		category := models.Category{
			Name:        "Test Category With Products",
			Description: "Category with products",
		}
		db.Create(&category)

		// Add products to the category
		product1 := models.Product{
			Name:        "Test Product in Category",
			Description: "Product in test category",
			Price:       19.99,
			Quantity:    5,
			CategoryID:  category.ID,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Another Test Product in Category",
			Description: "Another product in test category",
			Price:       29.99,
			Quantity:    7,
			CategoryID:  category.ID,
		}
		db.Create(&product2)

		// Test WithProducts scope
		var resultCategory models.Category
		result := db.Scopes(models.WithProducts).First(&resultCategory, category.ID)
		assert.Nil(t, result.Error)
		assert.Len(t, resultCategory.Products, 2)

		// Verify product details
		productNames := make(map[string]bool)
		for _, p := range resultCategory.Products {
			productNames[p.Name] = true
		}

		assert.True(t, productNames["Test Product in Category"])
		assert.True(t, productNames["Another Test Product in Category"])

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
		db.Unscoped().Delete(&category)
	})
}
