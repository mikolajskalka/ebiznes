package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Carts    []Cart `json:"carts" gorm:"foreignKey:UserID"`
}

// GORM Scopes
func ActiveUsers(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

func WithCarts(db *gorm.DB) *gorm.DB {
	return db.Preload("Carts")
}

func ByEmail(email string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email = ?", email)
	}
}
