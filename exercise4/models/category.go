package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products" gorm:"foreignKey:CategoryID"`
}

// GORM Scopes
func ActiveCategories(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

func WithProducts(db *gorm.DB) *gorm.DB {
	return db.Preload("Products")
}

func ByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+name+"%")
	}
}
