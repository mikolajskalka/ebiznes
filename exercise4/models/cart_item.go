package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// GORM Scopes
func ByCartID(cartID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("cart_id = ?", cartID)
	}
}

func WithProduct(db *gorm.DB) *gorm.DB {
	return db.Preload("Product")
}
