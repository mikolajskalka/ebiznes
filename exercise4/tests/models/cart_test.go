package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

func TestCartModel(t *testing.T) {
	const (
		userIDQuery     = "user_id = ?"
		cartIDQuery     = "cart_id = ?"
		testProductName = "Test Product for Cart"
	)

	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test data
	db.Unscoped().Where(userIDQuery, 999).Delete(&models.Cart{})
	db.Unscoped().Where("name = ?", testProductName).Delete(&models.Product{})

	// Create a test user ID
	testUserID := uint(999)

	// Create a test product for our cart items
	testProduct := models.Product{
		Name:        testProductName,
		Description: "Test product for cart tests",
		Price:       49.99,
		Quantity:    100,
		CategoryID:  1,
	}
	db.Create(&testProduct)

	// Test 1: Create a new cart
	t.Run("Create Cart", func(t *testing.T) {
		cart := models.Cart{
			UserID:     testUserID,
			TotalPrice: 0,
		}

		result := db.Create(&cart)
		assert.Nil(t, result.Error)
		assert.NotZero(t, cart.ID)
		assert.Equal(t, testUserID, cart.UserID)
		assert.Equal(t, 0.0, cart.TotalPrice)
	})

	// Test 2: Find a cart
	var savedCart models.Cart
	t.Run("Find Cart", func(t *testing.T) {
		result := db.Where(userIDQuery, testUserID).First(&savedCart)
		assert.Nil(t, result.Error)
		assert.Equal(t, testUserID, savedCart.UserID)
		assert.Equal(t, 0.0, savedCart.TotalPrice)
	})

	// Test 3: Add items to a cart
	t.Run("Add Cart Items", func(t *testing.T) {
		cartItem1 := models.CartItem{
			CartID:    savedCart.ID,
			ProductID: testProduct.ID,
			Quantity:  2,
			Price:     testProduct.Price,
		}

		result := db.Create(&cartItem1)
		assert.Nil(t, result.Error)
		assert.NotZero(t, cartItem1.ID)

		// Update cart total
		savedCart.TotalPrice = testProduct.Price * float64(cartItem1.Quantity)
		db.Save(&savedCart)

		// Verify cart total was updated
		var updatedCart models.Cart
		db.First(&updatedCart, savedCart.ID)
		assert.Equal(t, testProduct.Price*float64(cartItem1.Quantity), updatedCart.TotalPrice)
	})

	// Test 4: Get cart with items
	t.Run("Get Cart With Items", func(t *testing.T) {
		var cart models.Cart
		result := db.Scopes(models.WithCartItems).First(&cart, savedCart.ID)
		assert.Nil(t, result.Error)

		// Cart should have 1 item
		assert.Equal(t, 1, len(cart.Items))

		if len(cart.Items) > 0 {
			assert.Equal(t, testProduct.ID, cart.Items[0].ProductID)
			assert.Equal(t, 2, cart.Items[0].Quantity)
			assert.Equal(t, testProduct.Price, cart.Items[0].Price)
		}
	})

	// Test 5: Update cart item quantity
	t.Run("Update Cart Item Quantity", func(t *testing.T) {
		// Get the cart item
		var cartItem models.CartItem
		db.Where(cartIDQuery, savedCart.ID).First(&cartItem)

		// Update quantity
		cartItem.Quantity = 5
		db.Save(&cartItem)

		// Update cart total
		var cartItems []models.CartItem
		db.Where(cartIDQuery, savedCart.ID).Find(&cartItems)

		var totalPrice float64
		for _, item := range cartItems {
			totalPrice += item.Price * float64(item.Quantity)
		}

		savedCart.TotalPrice = totalPrice
		db.Save(&savedCart)

		// Verify update
		var updatedCart models.Cart
		db.First(&updatedCart, savedCart.ID)
		assert.Equal(t, testProduct.Price*5, updatedCart.TotalPrice)
	})

	// Test 6: Remove an item from cart
	t.Run("Remove Cart Item", func(t *testing.T) {
		// Get the cart item
		var cartItem models.CartItem
		db.Where(cartIDQuery, savedCart.ID).First(&cartItem)

		// Delete cart item
		db.Delete(&cartItem)

		// Update cart total price (should be 0)
		savedCart.TotalPrice = 0
		db.Save(&savedCart)

		// Verify cart item was removed
		var count int64
		db.Model(&models.CartItem{}).Where(cartIDQuery+" AND deleted_at IS NULL", savedCart.ID).Count(&count)
		assert.Equal(t, int64(0), count)

		// Verify cart total was updated
		var updatedCart models.Cart
		db.First(&updatedCart, savedCart.ID)
		assert.Equal(t, 0.0, updatedCart.TotalPrice)
	})

	// Test 7: Test ActiveCarts scope
	t.Run("ActiveCarts Scope", func(t *testing.T) {
		// Create a second cart
		cart2 := models.Cart{
			UserID:     testUserID,
			TotalPrice: 0,
		}
		db.Create(&cart2)

		// Delete one cart
		db.Delete(&cart2)

		// Test ActiveCarts scope
		var carts []models.Cart
		db.Scopes(models.ActiveCarts).Where(userIDQuery, testUserID).Find(&carts)

		// Should only find non-deleted carts (just 1)
		assert.Equal(t, 1, len(carts))
		assert.Equal(t, savedCart.ID, carts[0].ID)

		// Clean up
		db.Unscoped().Delete(&cart2)
	})

	// Test 8: Test ByUserID scope
	t.Run("ByUserID Scope", func(t *testing.T) {
		// Create a cart for a different user
		otherUserID := uint(888)
		otherCart := models.Cart{
			UserID:     otherUserID,
			TotalPrice: 0,
		}
		db.Create(&otherCart)

		// Test ByUserID scope
		var carts []models.Cart
		db.Scopes(models.ByUserID(testUserID)).Find(&carts)

		// Should only find carts with the test user ID
		foundTestUserCart := false
		foundOtherUserCart := false

		for _, c := range carts {
			if c.ID == savedCart.ID {
				foundTestUserCart = true
			}
			if c.ID == otherCart.ID {
				foundOtherUserCart = true
			}
		}

		assert.True(t, foundTestUserCart)
		assert.False(t, foundOtherUserCart)

		// Clean up
		db.Unscoped().Delete(&otherCart)
	})

	// Clean up all test data
	db.Unscoped().Where(userIDQuery, testUserID).Delete(&models.Cart{})
	db.Unscoped().Where("name = ?", testProductName).Delete(&models.Product{})
}
