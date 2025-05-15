package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

func TestCartItemModel(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test data
	db.Unscoped().Where("user_id = ?", 999).Delete(&models.Cart{})
	db.Unscoped().Where("name = ?", "Test Product for CartItem").Delete(&models.Product{})

	// Set up test data
	testUserID := uint(999)

	// Create a test cart
	testCart := models.Cart{
		UserID:     testUserID,
		TotalPrice: 0,
	}
	db.Create(&testCart)

	// Create a test product
	testProduct := models.Product{
		Name:        "Test Product for CartItem",
		Description: "Test product for cart item tests",
		Price:       29.99,
		Quantity:    50,
		CategoryID:  1,
	}
	db.Create(&testProduct)

	// Test 1: Create a new cart item
	var savedCartItem models.CartItem
	t.Run("Create CartItem", func(t *testing.T) {
		cartItem := models.CartItem{
			CartID:    testCart.ID,
			ProductID: testProduct.ID,
			Quantity:  3,
			Price:     testProduct.Price,
		}

		result := db.Create(&cartItem)
		assert.Nil(t, result.Error)
		assert.NotZero(t, cartItem.ID)
		assert.Equal(t, testCart.ID, cartItem.CartID)
		assert.Equal(t, testProduct.ID, cartItem.ProductID)
		assert.Equal(t, 3, cartItem.Quantity)
		assert.Equal(t, 29.99, cartItem.Price)

		savedCartItem = cartItem
	})

	// Test 2: Find a cart item
	t.Run("Find CartItem", func(t *testing.T) {
		var cartItem models.CartItem
		result := db.First(&cartItem, savedCartItem.ID)
		assert.Nil(t, result.Error)
		assert.Equal(t, savedCartItem.ID, cartItem.ID)
		assert.Equal(t, testCart.ID, cartItem.CartID)
		assert.Equal(t, testProduct.ID, cartItem.ProductID)
		assert.Equal(t, 3, cartItem.Quantity)
	})

	// Test 3: Update a cart item
	t.Run("Update CartItem", func(t *testing.T) {
		var cartItem models.CartItem
		result := db.First(&cartItem, savedCartItem.ID)
		assert.Nil(t, result.Error)

		cartItem.Quantity = 5
		db.Save(&cartItem)

		var updatedCartItem models.CartItem
		db.First(&updatedCartItem, savedCartItem.ID)
		assert.Equal(t, 5, updatedCartItem.Quantity)
	})

	// Test 4: Delete a cart item
	t.Run("Delete CartItem", func(t *testing.T) {
		var cartItem models.CartItem
		result := db.First(&cartItem, savedCartItem.ID)
		assert.Nil(t, result.Error)

		db.Delete(&cartItem)

		var deletedCartItem models.CartItem
		result = db.First(&deletedCartItem, savedCartItem.ID)
		assert.Error(t, result.Error) // Should not find the deleted item

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().First(&deletedCartItem, savedCartItem.ID)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedCartItem.DeletedAt.Time)
	})

	// Test 5: Test ByCartID scope
	t.Run("ByCartID Scope", func(t *testing.T) {
		// Create a new test cart
		testCart2 := models.Cart{
			UserID:     testUserID,
			TotalPrice: 0,
		}
		db.Create(&testCart2)

		// Create cart items for different carts
		cartItem1 := models.CartItem{
			CartID:    testCart.ID,
			ProductID: testProduct.ID,
			Quantity:  1,
			Price:     testProduct.Price,
		}
		db.Create(&cartItem1)

		cartItem2 := models.CartItem{
			CartID:    testCart2.ID,
			ProductID: testProduct.ID,
			Quantity:  2,
			Price:     testProduct.Price,
		}
		db.Create(&cartItem2)

		// Test ByCartID scope
		var cartItems []models.CartItem
		db.Unscoped().Scopes(models.ByCartID(testCart.ID)).Find(&cartItems)

		// Should only find items for the first cart
		cart1Found := false
		cart2Found := false

		for _, item := range cartItems {
			if item.ID == cartItem1.ID {
				cart1Found = true
			}
			if item.ID == cartItem2.ID {
				cart2Found = true
			}
		}

		assert.True(t, cart1Found)
		assert.False(t, cart2Found)

		// Clean up
		db.Unscoped().Delete(&cartItem1)
		db.Unscoped().Delete(&cartItem2)
		db.Unscoped().Delete(&testCart2)
	})

	// Test 6: Test WithProduct scope
	t.Run("WithProduct Scope", func(t *testing.T) {
		// Create a new cart item
		cartItem := models.CartItem{
			CartID:    testCart.ID,
			ProductID: testProduct.ID,
			Quantity:  4,
			Price:     testProduct.Price,
		}
		db.Create(&cartItem)

		// Test WithProduct scope
		var loadedCartItem models.CartItem
		result := db.Scopes(models.WithProduct).First(&loadedCartItem, cartItem.ID)
		assert.Nil(t, result.Error)

		// Should have product information loaded
		assert.Equal(t, testProduct.ID, loadedCartItem.Product.ID)
		assert.Equal(t, "Test Product for CartItem", loadedCartItem.Product.Name)
		assert.Equal(t, 29.99, loadedCartItem.Product.Price)

		// Clean up
		db.Unscoped().Delete(&cartItem)
	})

	// Clean up all test data
	db.Unscoped().Delete(&testProduct)
	db.Unscoped().Delete(&testCart)
}
