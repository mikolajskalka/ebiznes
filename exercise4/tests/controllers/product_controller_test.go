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

func setupEchoTest() (*echo.Echo, *httptest.ResponseRecorder) {
	e := echo.New()
	rec := httptest.NewRecorder()
	return e, rec
}

func TestProductController(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up test data to ensure a fresh test environment
	db.Unscoped().Where("name LIKE ?", "Test Controller Product%").Delete(&models.Product{})

	// Test data: category for product creation
	var testCategory models.Category
	result := db.First(&testCategory, 1)
	if result.Error != nil {
		t.Log("Creating test category since none exists")
		testCategory = models.Category{
			Name:        "Test Category",
			Description: "Test Category Description",
		}
		db.Create(&testCategory)
	}

	// Test 1: GetAllProducts
	t.Run("GetAllProducts", func(t *testing.T) {
		// Create test products to ensure we have data
		product1 := models.Product{
			Name:        "Test Controller Product 1",
			Description: "Test Description 1",
			Price:       99.99,
			Quantity:    10,
			CategoryID:  testCategory.ID,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Test Controller Product 2",
			Description: "Test Description 2",
			Price:       149.99,
			Quantity:    5,
			CategoryID:  testCategory.ID,
		}
		db.Create(&product2)

		// Setup test context
		e, rec := setupEchoTest()
		c := e.NewContext(httptest.NewRequest(http.MethodGet, "/products", nil), rec)

		// Call controller function
		if assert.NoError(t, controllers.GetAllProducts(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Product 1")
			assert.Contains(t, rec.Body.String(), "Test Controller Product 2")
		}
	})

	// Test 2: GetProduct
	var savedProductID uint
	t.Run("GetProduct", func(t *testing.T) {
		// Create a test product
		product := models.Product{
			Name:        "Test Controller Product Detail",
			Description: "Test Description Detail",
			Price:       199.99,
			Quantity:    8,
			CategoryID:  testCategory.ID,
		}
		db.Create(&product)
		savedProductID = product.ID

		// Setup test context
		e, rec := setupEchoTest()
		c := e.NewContext(httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", savedProductID), nil), rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", savedProductID))

		// Call controller function
		if assert.NoError(t, controllers.GetProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Product Detail")
			assert.Contains(t, rec.Body.String(), "199.99")
		}
	})

	// Test 3: GetProduct - Not Found
	t.Run("GetProduct_NotFound", func(t *testing.T) {
		e, rec := setupEchoTest()
		c := e.NewContext(httptest.NewRequest(http.MethodGet, "/products/9999", nil), rec)
		c.SetParamNames("id")
		c.SetParamValues("9999")

		// Call controller function
		if assert.NoError(t, controllers.GetProduct(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Contains(t, rec.Body.String(), "Product not found")
		}
	})

	// Test 4: CreateProduct
	t.Run("CreateProduct", func(t *testing.T) {
		jsonBody := `{
			"name": "Test Controller Product New",
			"description": "Test Description New",
			"price": 299.99,
			"quantity": 15,
			"category_id": 1
		}`

		e, rec := setupEchoTest()
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.CreateProduct(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Product New")
			assert.Contains(t, rec.Body.String(), "299.99")
		}

		// Verify product was created in database
		var product models.Product
		db.Where("name = ?", "Test Controller Product New").First(&product)
		assert.Equal(t, "Test Controller Product New", product.Name)
		assert.Equal(t, 299.99, product.Price)
	})

	// Test 5: UpdateProduct
	t.Run("UpdateProduct", func(t *testing.T) {
		// Create a product to update
		product := models.Product{
			Name:        "Test Controller Product To Update",
			Description: "Before Update",
			Price:       99.99,
			Quantity:    5,
			CategoryID:  testCategory.ID,
		}
		db.Create(&product)

		jsonBody := `{
			"name": "Test Controller Product Updated",
			"description": "After Update",
			"price": 129.99,
			"quantity": 10,
			"category_id": 1
		}`

		e, rec := setupEchoTest()
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d", product.ID), strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", product.ID))

		// Call controller function
		if assert.NoError(t, controllers.UpdateProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Product Updated")
			assert.Contains(t, rec.Body.String(), "After Update")
			assert.Contains(t, rec.Body.String(), "129.99")
		}

		// Verify product was updated in database
		var updatedProduct models.Product
		db.First(&updatedProduct, product.ID)
		assert.Equal(t, "Test Controller Product Updated", updatedProduct.Name)
		assert.Equal(t, "After Update", updatedProduct.Description)
		assert.Equal(t, 129.99, updatedProduct.Price)
		assert.Equal(t, 10, updatedProduct.Quantity)
	})

	// Test 6: DeleteProduct
	t.Run("DeleteProduct", func(t *testing.T) {
		// Create a product to delete
		product := models.Product{
			Name:        "Test Controller Product To Delete",
			Description: "Will be deleted",
			Price:       79.99,
			Quantity:    3,
			CategoryID:  testCategory.ID,
		}
		db.Create(&product)

		e, rec := setupEchoTest()
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", product.ID), nil)
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", product.ID))

		// Call controller function
		if assert.NoError(t, controllers.DeleteProduct(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}

		// Verify product was deleted (soft delete)
		var deletedProduct models.Product
		result := db.First(&deletedProduct, product.ID)
		assert.Error(t, result.Error) // Should not find the deleted product

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().First(&deletedProduct, product.ID)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedProduct.DeletedAt.Time)
	})

	// Clean up
	db.Unscoped().Where("name LIKE ?", "Test Controller Product%").Delete(&models.Product{})
	if testCategory.Name == "Test Category" {
		db.Unscoped().Delete(&testCategory)
	}
}
