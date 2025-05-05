package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}

// GORM Scopes
// Scope for active products
func ActiveProducts(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

// Scope for products with stock
func InStock(db *gorm.DB) *gorm.DB {
	return db.Where("quantity > 0")
}

// Scope for products by category
func ByCategoryID(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}

// Scope for products by price range
func ByPriceRange(min, max float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price BETWEEN ? AND ?", min, max)
	}
}
