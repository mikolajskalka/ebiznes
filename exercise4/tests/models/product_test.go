package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Test constants
const (
	// Query strings
	nameLikeQuery  = "name LIKE ?"
	nameEqualQuery = "name = ?"

	// Test patterns
	testProductPattern = "Test Product%"
	testCategoryPrefix = " Cat%"

	// Test product names
	testProduct1          = "Test Product 1"
	testProductActive     = "Test Product Active"
	testProductToDelete   = "Test Product To Delete"
	testProductInStock    = "Test Product In Stock"
	testProductOutOfStock = "Test Product Out of Stock"
	testProductCat1       = "Test Product Cat 1"
	testProductCat2       = "Test Product Cat 2"
	testProductLowPrice   = "Test Product Low Price"
	testProductMidPrice   = "Test Product Mid Price"
	testProductHighPrice  = "Test Product High Price"

	// Test descriptions
	descActive       = "Active product"
	descToDelete     = "Will be deleted"
	descInStock      = "Has stock"
	descOutOfStock   = "No stock"
	descCategory1    = "Category 1"
	descCategory2    = "Category 2"
	descLowPrice     = "Low price range"
	descMidPrice     = "Medium price range"
	descHighPrice    = "High price range"
	descTestProduct1 = "Test Description 1"
	descUpdated      = "Updated Description"

	// Price constants
	priceActive     = 19.99
	priceToDelete   = 29.99
	priceInStock    = 39.99
	priceOutOfStock = 49.99
	priceCat1       = 59.99
	priceCat2       = 69.99
	priceLow        = 10.99
	priceMid        = 50.99
	priceHigh       = 100.99
	priceInitial    = 99.99
	priceUpdated    = 129.99

	// Price range bounds
	minPriceRange = 30.0
	maxPriceRange = 80.0

	// Inventory quantities
	qtyDefault    = 5
	qtyInStock    = 10
	qtyOutOfStock = 0
	qtyCat1       = 8
	qtyCat2       = 6
	qtyBasicTest  = 10

	// Category IDs
	categoryID1 = uint(1)
	categoryID2 = uint(2)
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
		product := createTestProduct(t, db, testProduct1, descTestProduct1, priceInitial, qtyBasicTest, categoryID1)
		assert.Equal(t, testProduct1, product.Name)
		assert.Equal(t, priceInitial, product.Price)
	})

	// Test 2: Find a product
	t.Run("Find Product", func(t *testing.T) {
		product := findProductByName(t, db, nameEqualQuery, testProduct1)
		assert.Equal(t, testProduct1, product.Name)
		assert.Equal(t, descTestProduct1, product.Description)
		assert.Equal(t, priceInitial, product.Price)
	})

	// Test 3: Update a product
	t.Run("Update Product", func(t *testing.T) {
		product := findProductByName(t, db, nameEqualQuery, testProduct1)

		product.Price = priceUpdated
		product.Description = descUpdated

		result := db.Save(&product)
		assert.Nil(t, result.Error)

		updatedProduct := findProductByName(t, db, nameEqualQuery, testProduct1)
		assert.Equal(t, priceUpdated, updatedProduct.Price)
		assert.Equal(t, descUpdated, updatedProduct.Description)
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
		product1 := createTestProduct(t, db, testProductActive, descActive, priceActive, qtyDefault, categoryID1)
		product2 := createTestProduct(t, db, testProductToDelete, descToDelete, priceToDelete, qtyDefault-2, categoryID1)

		// Delete one product
		db.Delete(&product2)

		// Test ActiveProducts scope
		var products []models.Product
		db.Scopes(models.ActiveProducts).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results using a map
		productStatus := map[string]bool{
			testProductActive:   false,
			testProductToDelete: false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus[testProductActive])
		assert.False(t, productStatus[testProductToDelete])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test InStock scope
func testInStockScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("InStock Scope", func(t *testing.T) {
		// Create test products
		product1 := createTestProduct(t, db, testProductInStock, descInStock, priceInStock, qtyInStock, categoryID1)
		product2 := createTestProduct(t, db, testProductOutOfStock, descOutOfStock, priceOutOfStock, qtyOutOfStock, categoryID1)

		// Test InStock scope
		var products []models.Product
		db.Scopes(models.InStock).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results
		productStatus := map[string]bool{
			testProductInStock:    false,
			testProductOutOfStock: false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus[testProductInStock])
		assert.False(t, productStatus[testProductOutOfStock])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test ByCategoryID scope
func testByCategoryIDScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("ByCategoryID Scope", func(t *testing.T) {
		// Create test products
		product1 := createTestProduct(t, db, testProductCat1, descCategory1, priceCat1, qtyCat1, categoryID1)
		product2 := createTestProduct(t, db, testProductCat2, descCategory2, priceCat2, qtyCat2, categoryID2)

		// Test ByCategoryID scope
		var products []models.Product
		db.Scopes(models.ByCategoryID(categoryID1)).Where(nameLikeQuery, testProductPattern+testCategoryPrefix).Find(&products)

		// Verify results
		productStatus := map[string]bool{
			testProductCat1: false,
			testProductCat2: false,
		}
		verifyProductFound(t, products, productStatus)

		assert.True(t, productStatus[testProductCat1])
		assert.False(t, productStatus[testProductCat2])

		// Clean up
		deleteTestProducts(db, product1, product2)
	})
}

// Test ByPriceRange scope
func testByPriceRangeScope(t *testing.T, db *gorm.DB, nameLikeQuery, testProductPattern string) {
	t.Run("ByPriceRange Scope", func(t *testing.T) {
		// Create test products with different price ranges
		product1 := createTestProduct(t, db, testProductLowPrice, descLowPrice, priceLow, qtyDefault, categoryID1)
		product2 := createTestProduct(t, db, testProductMidPrice, descMidPrice, priceMid, qtyDefault, categoryID1)
		product3 := createTestProduct(t, db, testProductHighPrice, descHighPrice, priceHigh, qtyDefault, categoryID1)

		// Test ByPriceRange scope using min and max constants
		var products []models.Product
		db.Scopes(models.ByPriceRange(minPriceRange, maxPriceRange)).Where(nameLikeQuery, testProductPattern).Find(&products)

		// Verify results
		productStatus := map[string]bool{
			testProductLowPrice:  false,
			testProductMidPrice:  false,
			testProductHighPrice: false,
		}
		verifyProductFound(t, products, productStatus)

		assert.False(t, productStatus[testProductLowPrice])
		assert.True(t, productStatus[testProductMidPrice])
		assert.False(t, productStatus[testProductHighPrice])

		// Clean up
		deleteTestProducts(db, product1, product2, product3)
	})
}
