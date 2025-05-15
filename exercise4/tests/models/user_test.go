package tests

import (
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up any test users
	db.Unscoped().Where("email LIKE ?", "test%@example.com").Delete(&models.User{})

	// Test 1: Create a new user
	var savedUser models.User
	t.Run("Create User", func(t *testing.T) {
		user := models.User{
			Username: "testuser",
			Email:    "test@example.com",
		}

		result := db.Create(&user)
		assert.Nil(t, result.Error)
		assert.NotZero(t, user.ID)
		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "test@example.com", user.Email)

		savedUser = user
	})

	// Test 2: Find a user
	t.Run("Find User", func(t *testing.T) {
		var user models.User
		result := db.First(&user, savedUser.ID)
		assert.Nil(t, result.Error)
		assert.Equal(t, savedUser.ID, user.ID)
		assert.Equal(t, "testuser", user.Username)
		assert.Equal(t, "test@example.com", user.Email)
	})

	// Test 3: Update a user
	t.Run("Update User", func(t *testing.T) {
		var user models.User
		result := db.First(&user, savedUser.ID)
		assert.Nil(t, result.Error)

		user.Username = "updateduser"
		user.Email = "updated@example.com"
		db.Save(&user)

		var updatedUser models.User
		db.First(&updatedUser, savedUser.ID)
		assert.Equal(t, "updateduser", updatedUser.Username)
		assert.Equal(t, "updated@example.com", updatedUser.Email)
	})

	// Test 4: Delete a user (soft delete)
	t.Run("Delete User", func(t *testing.T) {
		var user models.User
		result := db.First(&user, savedUser.ID)
		assert.Nil(t, result.Error)

		db.Delete(&user)

		var deletedUser models.User
		result = db.First(&deletedUser, savedUser.ID)
		assert.Error(t, result.Error) // Should not find the deleted user

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().First(&deletedUser, savedUser.ID)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedUser.DeletedAt.Time)
	})

	// Test 5: Test ActiveUsers scope
	t.Run("ActiveUsers Scope", func(t *testing.T) {
		// Create two users
		user1 := models.User{
			Username: "activeuser",
			Email:    "test1@example.com",
		}
		db.Create(&user1)

		user2 := models.User{
			Username: "inactiveuser",
			Email:    "test2@example.com",
		}
		db.Create(&user2)

		// Delete one user
		db.Delete(&user2)

		// Test ActiveUsers scope
		var users []models.User
		db.Scopes(models.ActiveUsers).Where("email LIKE ?", "test%@example.com").Find(&users)

		// Should only find non-deleted users
		activeFound := false
		inactiveFound := false

		for _, u := range users {
			if u.Email == "test1@example.com" {
				activeFound = true
			}
			if u.Email == "test2@example.com" {
				inactiveFound = true
			}
		}

		assert.True(t, activeFound)
		assert.False(t, inactiveFound)

		// Clean up
		db.Unscoped().Delete(&user1)
		db.Unscoped().Delete(&user2)
	})

	// Test 6: Test ByEmail scope
	t.Run("ByEmail Scope", func(t *testing.T) {
		// Create users with different emails
		user1 := models.User{
			Username: "user1",
			Email:    "test3@example.com",
		}
		db.Create(&user1)

		user2 := models.User{
			Username: "user2",
			Email:    "test4@example.com",
		}
		db.Create(&user2)

		// Test ByEmail scope
		var users []models.User
		db.Scopes(models.ByEmail("test3@example.com")).Find(&users)

		// Should only find users with the matching email
		assert.Equal(t, 1, len(users))
		if len(users) > 0 {
			assert.Equal(t, "test3@example.com", users[0].Email)
		}

		// Clean up
		db.Unscoped().Delete(&user1)
		db.Unscoped().Delete(&user2)
	})

	// Test 7: User with Carts (WithCarts scope)
	t.Run("WithCarts Scope", func(t *testing.T) {
		// Create a test user
		user := models.User{
			Username: "userwithcarts",
			Email:    "test5@example.com",
		}
		db.Create(&user)

		// Create carts for the user
		cart1 := models.Cart{
			UserID:     user.ID,
			TotalPrice: 19.99,
		}
		db.Create(&cart1)

		cart2 := models.Cart{
			UserID:     user.ID,
			TotalPrice: 29.99,
		}
		db.Create(&cart2)

		// Test WithCarts scope
		var loadedUser models.User
		result := db.Scopes(models.WithCarts).First(&loadedUser, user.ID)
		assert.Nil(t, result.Error)

		// User should have 2 carts
		assert.Equal(t, 2, len(loadedUser.Carts))

		// Verify cart total prices
		cartTotals := make(map[float64]bool)
		for _, cart := range loadedUser.Carts {
			cartTotals[cart.TotalPrice] = true
		}

		assert.True(t, cartTotals[19.99])
		assert.True(t, cartTotals[29.99])

		// Clean up
		db.Unscoped().Where("user_id = ?", user.ID).Delete(&models.Cart{})
		db.Unscoped().Delete(&user)
	})
}
