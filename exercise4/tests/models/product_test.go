package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createTestProduct(t *testing.T, db *gorm.DB, name, description string, price float64, quantity int, categoryID uint) models.Product {
	product := models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CategoryID:  categoryID,
	}
	result := db.Create(&product)
	assert.Nil(t, result.Error)
	assert.NotZero(t, product.ID)
	return product
}

func deleteTestProducts(db *gorm.DB, products ...models.Product) {
	for _, product := range products {
		db.Unscoped().Delete(&product)
	}
}

func findProductByName(t *testing.T, db *gorm.DB, nameEqualQuery, name string) models.Product {
	var product models.Product
	result := db.Where(nameEqualQuery, name).First(&product)
	assert.Nil(t, result.Error)
	return product
}

func verifyProductFound(t *testing.T, products []models.Product, productMap map[string]bool) {
	for _, p := range products {
		if _, exists := productMap[p.Name]; exists {
			productMap[p.Name] = true
		}
	}
}

func TestProductModel(t *testing.T) {
	const (
		testProductPattern = "Test Product%"
		nameLikeQuery      = "name LIKE ?"
		nameEqualQuery     = "name = ?"
		testProduct1       = "Test Product 1"
	)

	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test products
	db.Unscoped().Where(nameLikeQuery, testProductPattern).Delete(&models.Product{})

	// Test basic CRUD operations
	testBasicCRUD(t, db, nameEqualQuery, testProduct1)

	// Test scopes
	testActiveProductsScope(t, db, nameLikeQuery, testProductPattern)
	testInStockScope(t, db, nameLikeQuery, testProductPattern)
	testByCategoryIDScope(t, db, nameLikeQuery, testProductPattern)
	testByPriceRangeScope(t, db, nameLikeQuery, testProductPattern)
}

// Test basic CRUD operations for Product model
func testBasicCRUD(t *testing.T, db *gorm.DB, nameEqualQuery, testProduct1 string) {
	// Test 1: Create a new product
	t.Run("Create Product", func(t *testing.T) {
		product := createTestProduct(t, db, testProduct1, "Test Description 1", 99.99, 10, 1)
		assert.Equal(t, testProduct1, product.Name)
		assert.Equal(t, 99.99, product.Price)
	})

	// Test 2: Find a product
	t.Run("Find Product", func(t *testing.T) {
		product := findProductByName(t, db, nameEqualQuery, testProduct1)
		assert.Equal(t, testProduct1, product.Name)
		assert.Equal(t, "Test Description 1", product.Description)
		assert.Equal(t, 99.99, product.Price)
	})

	// Test 3: Update a product
	t.Run("Update Product", func(t *testing.T) {
		product := findProductByName(t, db, nameEqualQuery, testProduct1)

		product.Price = 129.99
		product.Description = "Updated Description"

		result := db.Save(&product)
		assert.Nil(t, result.Error)

		updatedProduct := findProductByName(t, db, nameEqualQuery, testProduct1)
		assert.Equal(t, 129.99, updatedProduct.Price)
		assert.Equal(t, "Updated Description", updatedProduct.Description)
	})

	// Test 4: Delete a product (soft delete)
	t.Run("Delete Product", func(t *testing.T) {
		product := findProductByName(t, db, nameEqualQuery, testProduct1)

		result := db.Delete(&product)
		assert.Nil(t, result.Error)

		// Should not find the deleted product
		var deletedProduct models.Product
		result = db.Where(nameEqualQuery, testProduct1).First(&deletedProduct)
		assert.Error(t, result.Error)

		// Should find with Unscoped
		result = db.Unscoped().Where(nameEqualQuery, testProduct1).First(&deletedProduct)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedProduct.DeletedAt.Time)
	})
}

// Test ActiveProducts scope
func testActiveProductsScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("ActiveProducts Scope", func(t *testing.T) {
		// Create test products
		product1 := createTestProduct(t, db, "Test Product Active", "Active product", 19.99, 5, 1)
		product2 := createTestProduct(t, db, "Test Product To Delete", "Will be deleted", 29.99, 3, 1)

		// Delete one product
		db.Delete(&product2)

		// Test ActiveProducts scope
		var products []models.Product
		db.Scopes(models.ActiveProducts).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results using a map
		productStatus := map[string]bool{
			"Test Product Active":    false,
			"Test Product To Delete": false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus["Test Product Active"])
		assert.False(t, productStatus["Test Product To Delete"])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test InStock scope
func testInStockScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("InStock Scope", func(t *testing.T) {
		// Create test products
		product1 := createTestProduct(t, db, "Test Product In Stock", "Has stock", 39.99, 10, 1)
		product2 := createTestProduct(t, db, "Test Product Out of Stock", "No stock", 49.99, 0, 1)

		// Test InStock scope
		var products []models.Product
		db.Scopes(models.InStock).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results
		productStatus := map[string]bool{
			"Test Product In Stock":     false,
			"Test Product Out of Stock": false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus["Test Product In Stock"])
		assert.False(t, productStatus["Test Product Out of Stock"])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test ByCategoryID scope
func testByCategoryIDScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("ByCategoryID Scope", func(t *testing.T) {
		// Create test products
		product1 := createTestProduct(t, db, "Test Product Cat 1", "Category 1", 59.99, 8, 1)
		product2 := createTestProduct(t, db, "Test Product Cat 2", "Category 2", 69.99, 6, 2)

		// Test ByCategoryID scope
		var products []models.Product
		db.Scopes(models.ByCategoryID(1)).Where(nameLikeQuery, testProductPattern+" Cat%").Find(&products)

		// Verify results
		productStatus := map[string]bool{
			"Test Product Cat 1": false,
			"Test Product Cat 2": false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus["Test Product Cat 1"])
		assert.False(t, productStatus["Test Product Cat 2"])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test ByPriceRange scope
func testByPriceRangeScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("ByPriceRange Scope", func(t *testing.T) {
		// Create test products with different price ranges
		product1 := createTestProduct(t, db, "Test Product Low Price", "Low price range", 10.99, 5, 1)
		product2 := createTestProduct(t, db, "Test Product Mid Price", "Medium price range", 50.99, 5, 1)
		product3 := createTestProduct(t, db, "Test Product High Price", "High price range", 100.99, 5, 1)

		// Test ByPriceRange scope (30-80)
		var products []models.Product
		db.Scopes(models.ByPriceRange(30.0, 80.0)).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results
		productStatus := map[string]bool{
			"Test Product Low Price":  false,
			"Test Product Mid Price":  false,
			"Test Product High Price": false,
		}
		verifyProductFound(t, products, productStatus)

		assert.False(t, productStatus["Test Product Low Price"])
		assert.True(t, productStatus["Test Product Mid Price"])
		assert.False(t, productStatus["Test Product High Price"])

		// Clean up
		deleteTestProducts(db, product1, product2, product3)
	})
}
