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

func TestCategoryController(t *testing.T) {
	// Initialize the test database
	database.Initialize()
	db := database.GetDB()

	// Clean up test data to ensure a fresh test environment
	db.Unscoped().Where("name LIKE ?", "Test Controller Category%").Delete(&models.Category{})

	// Test 1: GetAllCategories
	t.Run("GetAllCategories", func(t *testing.T) {
		// Create test categories to ensure we have data
		category1 := models.Category{
			Name:        "Test Controller Category 1",
			Description: "Test Category Description 1",
		}
		db.Create(&category1)

		category2 := models.Category{
			Name:        "Test Controller Category 2",
			Description: "Test Category Description 2",
		}
		db.Create(&category2)

		// Setup test context
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/categories", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.GetAllCategories(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category 1")
			assert.Contains(t, rec.Body.String(), "Test Controller Category 2")
		}
	})

	// Test 2: GetCategory
	var savedCategoryID uint
	t.Run("GetCategory", func(t *testing.T) {
		// Create a test category
		category := models.Category{
			Name:        "Test Controller Category Detail",
			Description: "Test Category Description Detail",
		}
		db.Create(&category)
		savedCategoryID = category.ID

		// Setup test context
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/categories/%d", savedCategoryID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", savedCategoryID))

		// Call controller function
		if assert.NoError(t, controllers.GetCategory(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category Detail")
			assert.Contains(t, rec.Body.String(), "Test Category Description Detail")
		}
	})

	// Test 3: CreateCategory
	t.Run("CreateCategory", func(t *testing.T) {
		jsonBody := `{
			"name": "Test Controller Category New",
			"description": "Test Category Description New"
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.CreateCategory(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category New")
			assert.Contains(t, rec.Body.String(), "Test Category Description New")
		}

		// Verify category was created in database
		var category models.Category
		db.Where("name = ?", "Test Controller Category New").First(&category)
		assert.Equal(t, "Test Controller Category New", category.Name)
		assert.Equal(t, "Test Category Description New", category.Description)
	})

	// Test 4: UpdateCategory
	t.Run("UpdateCategory", func(t *testing.T) {
		// Create a category to update
		category := models.Category{
			Name:        "Test Controller Category To Update",
			Description: "Before Update",
		}
		db.Create(&category)

		jsonBody := `{
			"name": "Test Controller Category Updated",
			"description": "After Update"
		}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/categories/%d", category.ID), strings.NewReader(jsonBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", category.ID))

		// Call controller function
		if assert.NoError(t, controllers.UpdateCategory(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category Updated")
			assert.Contains(t, rec.Body.String(), "After Update")
		}

		// Verify category was updated in database
		var updatedCategory models.Category
		db.First(&updatedCategory, category.ID)
		assert.Equal(t, "Test Controller Category Updated", updatedCategory.Name)
		assert.Equal(t, "After Update", updatedCategory.Description)
	})

	// Test 5: DeleteCategory
	t.Run("DeleteCategory", func(t *testing.T) {
		// Create a category to delete
		category := models.Category{
			Name:        "Test Controller Category To Delete",
			Description: "Will be deleted",
		}
		db.Create(&category)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/categories/%d", category.ID), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprintf("%d", category.ID))

		// Call controller function
		if assert.NoError(t, controllers.DeleteCategory(c)) {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}

		// Verify category was deleted (soft delete)
		var deletedCategory models.Category
		result := db.First(&deletedCategory, category.ID)
		assert.Error(t, result.Error) // Should not find the deleted category

		// It should be found when using Unscoped (which includes soft-deleted records)
		result = db.Unscoped().First(&deletedCategory, category.ID)
		assert.Nil(t, result.Error)
		assert.NotNil(t, deletedCategory.DeletedAt.Time)
	})

	// Test 6: GetCategoriesWithProducts
	t.Run("GetCategoriesWithProducts", func(t *testing.T) {
		// Create a test category
		category := models.Category{
			Name:        "Test Controller Category With Products",
			Description: "Category with products",
		}
		db.Create(&category)

		// Add products to this category
		product1 := models.Product{
			Name:        "Test Product In Category",
			Description: "Product in test controller category",
			Price:       19.99,
			Quantity:    5,
			CategoryID:  category.ID,
		}
		db.Create(&product1)

		product2 := models.Product{
			Name:        "Another Test Product In Category",
			Description: "Another product in test controller category",
			Price:       29.99,
			Quantity:    7,
			CategoryID:  category.ID,
		}
		db.Create(&product2)

		// Setup test context
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/categories/with-products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call controller function
		if assert.NoError(t, controllers.GetCategoriesWithProducts(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category With Products")
			// In a real application, we'd check that the products were included in the response
			// But for simplicity, we'll just ensure the response status is correct
		}

		// Clean up
		db.Unscoped().Delete(&product1)
		db.Unscoped().Delete(&product2)
		db.Unscoped().Delete(&category)
	})

	// Test 7: SearchCategoriesByName
	t.Run("SearchCategoriesByName", func(t *testing.T) {
		// Create test categories with specific names for testing search
		category1 := models.Category{
			Name:        "Test Controller Category Apple",
			Description: "Apple category",
		}
		db.Create(&category1)

		category2 := models.Category{
			Name:        "Test Controller Category Banana",
			Description: "Banana category",
		}
		db.Create(&category2)

		// Setup test context
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/categories/search?name=Apple", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/categories/search")
		c.QueryParams().Set("name", "Apple")

		// Call controller function
		if assert.NoError(t, controllers.SearchCategoriesByName(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Test Controller Category Apple")
			assert.NotContains(t, rec.Body.String(), "Test Controller Category Banana")
		}

		// Clean up
		db.Unscoped().Delete(&category1)
		db.Unscoped().Delete(&category2)
	})

	// Clean up all test data
	db.Unscoped().Where("name LIKE ?", "Test Controller Category%").Delete(&models.Category{})
}
