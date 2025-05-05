package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	Items      []CartItem `json:"items" gorm:"foreignKey:CartID"`
	TotalPrice float64    `json:"total_price"`
}

// GORM Scopes
func ActiveCarts(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

func WithCartItems(db *gorm.DB) *gorm.DB {
	return db.Preload("Items")
}

func ByUserID(userID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID)
	}
}
