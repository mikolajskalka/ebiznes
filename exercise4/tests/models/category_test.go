package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Helper functions to reduce cognitive complexity
func createTestCategory(t *testing.T, db *gorm.DB, name, description string) models.Category {
	category := models.Category{
		Name:        name,
		Description: description,
	}
	result := db.Create(&category)
	assert.Nil(t, result.Error)
	assert.NotZero(t, category.ID)
	return category
}

func cleanupCategories(db *gorm.DB, categories ...models.Category) {
	for _, category := range categories {
		db.Unscoped().Delete(&category)
	}
}

func cleanupProducts(db *gorm.DB, products ...models.Product) {
	for _, product := range products {
		db.Unscoped().Delete(&product)
	}
}

func TestCategoryModel(t *testing.T) {
	const (
		testCategoryPattern      = "Test Category%"
		nameLikeQuery            = "name LIKE ?"
		nameEqualQuery           = "name = ?"
		testCategory1            = "Test Category 1"
		testCategoryDescription1 = "Test Category Description 1"
		updatedTestCategory      = "Updated Test Category"
	)

	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test categories
	db.Unscoped().Where(nameLikeQuery, testCategoryPattern).Delete(&models.Category{})

	// Test 1: Create a new category
	t.Run("Create Category", func(t *testing.T) {
		category := createTestCategory(t, db, testCategory1, testCategoryDescription1)
		assert.Equal(t, testCategory1, category.Name)
		assert.Equal(t, testCategoryDescription1, category.Description)
	})

	// Test 2: Find a category
	t.Run("Find Category", func(t *testing.T) {
		var category models.Category
		result := db.Where(nameEqualQuery, testCategory1).First(&category)
		assert.Nil(t, result.Error)
		assert.Equal(t, testCategory1, category.Name)
		assert.Equal(t, testCategoryDescription1, category.Description)
	})

	// Test 3: Update a category
	t.Run("Update Category", func(t *testing.T) {
		var category models.Category
		db.Where(nameEqualQuery, testCategory1).First(&category)

		category.Name = updatedTestCategory
		category.Description = "Updated Category Description"

		result := db.Save(&category)
		assert.Nil(t, result.Error)

		var updatedCategory models.Category
		db.First(&updatedCategory, category.ID)
		assert.Equal(t, updatedTestCategory, updatedCategory.Name)
		assert.Equal(t, "Updated Category Description", updatedCategory.Description)
	})

	// Test 4: Delete a category (soft delete)
	t.Run("Delete Category", func(t *testing.T) {
		var category models.Category
		db.Where(nameEqualQuery, updatedTestCategory).First(&category)

		result := db.Delete(&category)
		assert.Nil(t, result.Error)

		var deletedCategory models.Category
		result = db.Where(nameEqualQuery, updatedTestCategory).First(&deletedCategory)
		assert.Error(t, result.Error) // Should not find the deleted category

		// Check it exists when using Unscoped
		result = db.Unscoped().Where(nameEqualQuery, updatedTestCategory).First(&deletedCategory)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedCategory.DeletedAt.Time)
	})

	// Test 5: Test ActiveCategories scope
	t.Run("ActiveCategories Scope", func(t *testing.T) {
		// Create categories
		category1 := createTestCategory(t, db, "Test Category Active", "Active category")
		category2 := createTestCategory(t, db, "Test Category To Delete", "Will be deleted")

		// Delete one category
		db.Delete(&category2)

		// Test ActiveCategories scope
		var categories []models.Category
		db.Scopes(models.ActiveCategories).Where(nameLikeQuery, testCategoryPattern).Find(&categories)

		// Verify results
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
		cleanupCategories(db, category1, category2)
	})

	// Test 6: Test ByName scope
	t.Run("ByName Scope", func(t *testing.T) {
		// Create test categories
		category1 := createTestCategory(t, db, "Test Category Apple", "Apple products")
		category2 := createTestCategory(t, db, "Test Category Banana", "Banana products")
		category3 := createTestCategory(t, db, "Test Category Apple Premium", "Premium apple products")

		testByNameScope(t, db, nameLikeQuery, testCategoryPattern)

		// Clean up
		cleanupCategories(db, category1, category2, category3)
	})

	// Test 7: Test WithProducts scope
	t.Run("WithProducts Scope", func(t *testing.T) {
		// Create a test category with products
		category := createTestCategory(t, db, "Test Category With Products", "Category with products")

		// Add products
		products := createProductsForCategory(t, db, category.ID)

		// Test WithProducts scope
		var resultCategory models.Category
		result := db.Scopes(models.WithProducts).First(&resultCategory, category.ID)
		assert.Nil(t, result.Error)
		assert.Len(t, resultCategory.Products, 2)

		// Verify product names
		productNames := make(map[string]bool)
		for _, p := range resultCategory.Products {
			productNames[p.Name] = true
		}
		assert.True(t, productNames["Test Product in Category"])
		assert.True(t, productNames["Another Test Product in Category"])

		// Clean up
		cleanupProducts(db, products...)
		cleanupCategories(db, category)
	})
}

func testByNameScope(t *testing.T, db *gorm.DB, nameLikeQuery, testCategoryPattern string) {
	// Test ByName scope
	var categories []models.Category
	db.Scopes(models.ByName("Apple")).Where(nameLikeQuery, testCategoryPattern).Find(&categories)

	// Check results
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
}

func createProductsForCategory(t *testing.T, db *gorm.DB, categoryID uint) []models.Product {
	product1 := models.Product{
		Name:        "Test Product in Category",
		Description: "Product in test category",
		Price:       19.99,
		Quantity:    5,
		CategoryID:  categoryID,
	}
	db.Create(&product1)

	product2 := models.Product{
		Name:        "Another Test Product in Category",
		Description: "Another product in test category",
		Price:       29.99,
		Quantity:    7,
		CategoryID:  categoryID,
	}
	db.Create(&product2)

	return []models.Product{product1, product2}
}
