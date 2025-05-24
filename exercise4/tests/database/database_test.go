package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Helper function to verify required table columns
func verifyTableColumns(t *testing.T, db *gorm.DB, tableName string, requiredColumns map[string]bool) {
	// Query table info
	type ColumnInfo struct {
		Name string
	}
	var columns []ColumnInfo
	result := db.Raw("PRAGMA table_info(" + tableName + ")").Scan(&columns)
	assert.Nil(t, result.Error)

	// Mark found columns
	for _, column := range columns {
		if _, exists := requiredColumns[column.Name]; exists {
			requiredColumns[column.Name] = true
		}
	}

	// Check if all required columns were found
	for column, found := range requiredColumns {
		assert.True(t, found, "Column "+column+" should exist in "+tableName+" table")
	}
}

func TestDatabaseInitialization(t *testing.T) {
	// Test 1: Database initialization
	t.Run("Initialize Database", func(t *testing.T) {
		// Initialize the database
		database.Initialize()

		// Get DB instance
		db := database.GetDB()

		// Ensure DB instance is not nil
		assert.NotNil(t, db)

		// Try a simple query to verify connection
		var count int64
		result := db.Raw("SELECT 1").Count(&count)
		assert.Nil(t, result.Error)
	})

	// Test 2: GetDB returns the same instance
	t.Run("GetDB Returns Same Instance", func(t *testing.T) {
		// Get DB instance twice
		db1 := database.GetDB()
		db2 := database.GetDB()

		// Both should be non-nil
		assert.NotNil(t, db1)
		assert.NotNil(t, db2)

		// They should be the same instance
		assert.Equal(t, db1, db2)
	})

	// Test 3: Database tables are created
	t.Run("Database Tables Created", func(t *testing.T) {
		db := database.GetDB()

		// Check if tables exist
		tables := []string{"users", "categories", "products", "carts", "cart_items"}

		for _, table := range tables {
			var count int64
			result := db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", table).Count(&count)
			assert.Nil(t, result.Error)
			assert.Greater(t, count, int64(0), "Table "+table+" should exist")
		}
	})

	// Test 4: Product table has required columns
	t.Run("Product Table Has Required Columns", func(t *testing.T) {
		db := database.GetDB()

		// Check required columns for products table
		productColumns := map[string]bool{
			"id":          false,
			"created_at":  false,
			"updated_at":  false,
			"deleted_at":  false,
			"name":        false,
			"description": false,
			"price":       false,
			"quantity":    false,
			"category_id": false,
		}

		verifyTableColumns(t, db, "products", productColumns)
	})

	// Test 5: Cart table has required columns
	t.Run("Cart Table Has Required Columns", func(t *testing.T) {
		db := database.GetDB()

		// Check required columns for carts table
		cartColumns := map[string]bool{
			"id":          false,
			"created_at":  false,
			"updated_at":  false,
			"deleted_at":  false,
			"user_id":     false,
			"total_price": false,
		}

		verifyTableColumns(t, db, "carts", cartColumns)
	})
}
