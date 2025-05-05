package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
)

// GetAllCarts - Get all carts
func GetAllCarts(c echo.Context) error {
	var carts []models.Cart

	db := database.GetDB()
	result := db.Scopes(models.ActiveCarts).Find(&carts)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve carts",
		})
	}

	return c.JSON(http.StatusOK, carts)
}

// GetCart - Get a cart by ID
func GetCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid cart ID",
		})
	}

	var cart models.Cart
	db := database.GetDB()
	result := db.Scopes(models.WithCartItems).First(&cart, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Cart not found",
		})
	}

	return c.JSON(http.StatusOK, cart)
}

// CreateCart - Create a new cart
func CreateCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid cart data",
		})
	}

	db := database.GetDB()
	result := db.Create(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create cart",
		})
	}

	return c.JSON(http.StatusCreated, cart)
}

// AddItemToCart - Add an item to cart
func AddItemToCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid cart ID",
		})
	}

	// Check if cart exists
	var cart models.Cart
	db := database.GetDB()
	result := db.First(&cart, cartID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Cart not found",
		})
	}

	// Bind cart item data
	cartItem := new(models.CartItem)
	if err := c.Bind(cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid cart item data",
		})
	}

	// Set cart ID
	cartItem.CartID = uint(cartID)

	// Get product details
	var product models.Product
	result = db.First(&product, cartItem.ProductID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	// Set price from product
	cartItem.Price = product.Price

	// Save cart item
	result = db.Create(&cartItem)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to add item to cart",
		})
	}

	// Update cart total price
	var cartItems []models.CartItem
	db.Where("cart_id = ?", cartID).Find(&cartItems)

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.Price * float64(item.Quantity)
	}

	cart.TotalPrice = totalPrice
	db.Save(&cart)

	return c.JSON(http.StatusCreated, cartItem)
}

// RemoveItemFromCart - Remove an item from cart
func RemoveItemFromCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid cart ID",
		})
	}

	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid item ID",
		})
	}

	// Check if cart exists
	var cart models.Cart
	db := database.GetDB()
	result := db.First(&cart, cartID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Cart not found",
		})
	}

	// Check if cart item exists
	var cartItem models.CartItem
	result = db.Where("cart_id = ? AND id = ?", cartID, itemID).First(&cartItem)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Cart item not found",
		})
	}

	// Delete cart item
	db.Delete(&cartItem)

	// Update cart total price
	var cartItems []models.CartItem
	db.Where("cart_id = ?", cartID).Find(&cartItems)

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.Price * float64(item.Quantity)
	}

	cart.TotalPrice = totalPrice
	db.Save(&cart)

	return c.NoContent(http.StatusNoContent)
}

// GetCartByUser - Get cart by user ID using scope
func GetCartByUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	var carts []models.Cart
	db := database.GetDB()
	result := db.Scopes(models.ActiveCarts, models.WithCartItems, models.ByUserID(uint(userID))).Find(&carts)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve carts",
		})
	}

	return c.JSON(http.StatusOK, carts)
}
