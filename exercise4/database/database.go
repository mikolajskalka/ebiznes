package database

import (
	"log"
	"os"

	"github.com/mikolajskalka/ebiznes/exercise4/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize database connection and migrate models
func Initialize() {
	var err error
	// Check if the data directory exists, if so use that path
	dbPath := "shop.db"
	if _, err := os.Stat("/app/data"); err == nil {
		dbPath = "/app/data/shop.db"
	}
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}

// Get DB instance
func GetDB() *gorm.DB {
	return DB
}
