package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
)

// GetAllCategories - Get all categories
func GetAllCategories(c echo.Context) error {
	var categories []models.Category

	db := database.GetDB()
	result := db.Scopes(models.ActiveCategories).Find(&categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

// GetCategoriesWithProducts - Get all categories with products
func GetCategoriesWithProducts(c echo.Context) error {
	var categories []models.Category

	db := database.GetDB()
	result := db.Scopes(models.ActiveCategories, models.WithProducts).Find(&categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve categories with products",
		})
	}

	return c.JSON(http.StatusOK, categories)
}

// GetCategory - Get a category by ID
func GetCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	var category models.Category
	db := database.GetDB()
	result := db.First(&category, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Category not found",
		})
	}

	return c.JSON(http.StatusOK, category)
}

// CreateCategory - Create a new category
func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category data",
		})
	}

	db := database.GetDB()
	result := db.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create category",
		})
	}

	return c.JSON(http.StatusCreated, category)
}

// UpdateCategory - Update an existing category
func UpdateCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	// Check if category exists
	var existingCategory models.Category
	db := database.GetDB()
	result := db.First(&existingCategory, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Category not found",
		})
	}

	// Bind updated data
	updatedCategory := new(models.Category)
	if err := c.Bind(updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category data",
		})
	}

	// Update category
	existingCategory.Name = updatedCategory.Name
	existingCategory.Description = updatedCategory.Description

	db.Save(&existingCategory)

	return c.JSON(http.StatusOK, existingCategory)
}

// DeleteCategory - Delete a category
func DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	// Check if category exists
	var category models.Category
	db := database.GetDB()
	result := db.First(&category, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Category not found",
		})
	}

	// Delete category (soft delete with GORM)
	db.Delete(&category)

	return c.NoContent(http.StatusNoContent)
}

// SearchCategoriesByName - Search categories by name using scope
func SearchCategoriesByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name parameter is required",
		})
	}

	var categories []models.Category
	db := database.GetDB()
	result := db.Scopes(models.ActiveCategories, models.ByName(name)).Find(&categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to search categories",
		})
	}

	return c.JSON(http.StatusOK, categories)
}
