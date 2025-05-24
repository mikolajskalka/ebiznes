package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikolajskalka/ebiznes/exercise4/controllers"
	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"github.com/stretchr/testify/assert"
)

const (
	userIDQuery     = "user_id = ?"
	testProductName = "Test Cart Controller Product"
)

func TestCartController(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up test data to ensure a fresh test environment
	db.Unscoped().Where(userIDQuery, 999).Delete(&models.Cart{})
	db.Unscoped().Where("name = ?", testProductName).Delete(&models.Product{})

	// Create test user ID
	testUserID := uint(999)

	// Create a test product for our cart items
	testProduct := models.Product{
		Name:        testProductName,
		Description: "Test product for cart controller tests",
		Price:       49.99,
		Quantity:    100,
		CategoryID:  1,
	}
	db.Create(&testProduct)

	// Test 1: CreateCart
	var savedCartID uint
	t.Run("CreateCart", func(t *testing.T) {
		jsonBody := `{
			"user_id": 999,
			"total_price": 0
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.CreateCart(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Contains(t, rec.Body.String(), `"user_id":999`)
			assert.Contains(t, rec.Body.String(), `"total_price":0`)
		}

		// Extract created cart ID for later tests
		var cart models.Cart
		db.Where(userIDQuery, testUserID).First(&cart)
		savedCartID = cart.ID
	})

	// Test 2: GetCart
	t.Run("GetCart", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/carts/1", nil) // Using a simple fixed path for testing
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", savedCartID))

		// Call controller function
		if assert.NoError(t, controllers.GetCart(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `"user_id":999`)
		}
	})

	// Test 3: GetAllCarts
	t.Run("GetAllCarts", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/carts", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.GetAllCarts(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// The response should contain our test cart with user_id 999
			assert.Contains(t, rec.Body.String(), `"user_id":999`)
		}
	})

	// Test 4: AddItemToCart
	t.Run("AddItemToCart", func(t *testing.T) {
		jsonBody := fmt.Sprintf(`{
			"product_id": %d,
			"quantity": 3
		}`, testProduct.ID)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/carts/%d/items", savedCartID), strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", savedCartID))

		// Call controller function
		if assert.NoError(t, controllers.AddItemToCart(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Contains(t, rec.Body.String(), `"product_id":`)
			assert.Contains(t, rec.Body.String(), `"quantity":3`)
		}

		// Verify cart total was updated
		var cart models.Cart
		db.First(&cart, savedCartID)
		assert.Equal(t, testProduct.Price*3, cart.TotalPrice)
	})

	// Test 5: GetCartByUser
	t.Run("GetCartByUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/carts/user/999", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("userId")
		c.SetParamValues("999")

		// Call controller function
		if assert.NoError(t, controllers.GetCartByUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `"user_id":999`)
			// Verify the response contains cart items
			assert.Contains(t, rec.Body.String(), `"items":[`)
		}
	})

	// Test 6: RemoveItemFromCart
	t.Run("RemoveItemFromCart", func(t *testing.T) {
		// Get the cart item ID
		var cartItem models.CartItem
		db.Where("cart_id = ?", savedCartID).First(&cartItem)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/carts/%d/items/%d", savedCartID, cartItem.ID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id", "itemId")
		c.SetParamValues(fmt.Sprintf("%d", savedCartID), fmt.Sprintf("%d", cartItem.ID))

		// Call controller function
		if assert.NoError(t, controllers.RemoveItemFromCart(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}

		// Verify cart item was removed
		var count int64
		db.Model(&models.CartItem{}).Where("id = ? AND deleted_at IS NULL", cartItem.ID).Count(&count)
		assert.Equal(t, int64(0), count)

		// Verify cart total was updated
		var cart models.Cart
		db.First(&cart, savedCartID)
		assert.Equal(t, 0.0, cart.TotalPrice)
	})

	// Clean up test data
	db.Unscoped().Where(userIDQuery, testUserID).Delete(&models.Cart{})
	db.Unscoped().Where("name = ?", testProductName).Delete(&models.Product{})
}
