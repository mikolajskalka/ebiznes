package main

import (
	"log"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
	"github.com/mikolajskalka/ebiznes/exercise4/models"
)

func main() {
	log.Println("Starting to seed products...")

	// Initialize database
	database.Initialize()

	// Get database instance
	db := database.GetDB()

	// Create categories if they don't exist
	categories := []models.Category{
		{Name: "Electronics", Description: "Electronic devices and gadgets"},
		{Name: "Clothing", Description: "Apparel and fashion items"},
		{Name: "Books", Description: "Books and literature"},
		{Name: "Home & Kitchen", Description: "Home and kitchen items"},
		{Name: "Sports", Description: "Sports equipment and gear"},
	}

	for _, category := range categories {
		var existingCategory models.Category
		result := db.Where("name = ?", category.Name).First(&existingCategory)
		if result.RowsAffected == 0 {
			db.Create(&category)
			log.Printf("Created category: %s\n", category.Name)
		} else {
			log.Printf("Category already exists: %s\n", category.Name)
		}
	}

	// Add products with their respective categories
	products := []models.Product{
		// Electronics
		{Name: "Smartphone", Description: "Latest model smartphone", Price: 999.99, Quantity: 50, CategoryID: 1},
		{Name: "Laptop", Description: "High-performance laptop", Price: 1299.99, Quantity: 30, CategoryID: 1},
		{Name: "Wireless Headphones", Description: "Noise-canceling wireless headphones", Price: 199.99, Quantity: 100, CategoryID: 1},
		{Name: "Tablet", Description: "10-inch tablet with retina display", Price: 499.99, Quantity: 45, CategoryID: 1},
		{Name: "Smartwatch", Description: "Fitness and health tracking smartwatch", Price: 249.99, Quantity: 75, CategoryID: 1},

		// Clothing
		{Name: "T-Shirt", Description: "Cotton crew neck t-shirt", Price: 19.99, Quantity: 200, CategoryID: 2},
		{Name: "Jeans", Description: "Classic blue denim jeans", Price: 59.99, Quantity: 150, CategoryID: 2},
		{Name: "Hoodie", Description: "Warm pullover hoodie", Price: 39.99, Quantity: 100, CategoryID: 2},
		{Name: "Sneakers", Description: "Comfortable everyday sneakers", Price: 79.99, Quantity: 80, CategoryID: 2},
		{Name: "Winter Jacket", Description: "Waterproof insulated winter jacket", Price: 129.99, Quantity: 60, CategoryID: 2},

		// Books
		{Name: "Programming in Go", Description: "Learn Go programming language", Price: 34.99, Quantity: 40, CategoryID: 3},
		{Name: "Science Fiction Anthology", Description: "Collection of sci-fi short stories", Price: 24.99, Quantity: 35, CategoryID: 3},
		{Name: "Cookbook", Description: "International cuisine recipes", Price: 29.99, Quantity: 25, CategoryID: 3},
		{Name: "History Book", Description: "Comprehensive world history", Price: 49.99, Quantity: 20, CategoryID: 3},
		{Name: "Self-Help Guide", Description: "Personal development and growth", Price: 19.99, Quantity: 50, CategoryID: 3},

		// Home & Kitchen
		{Name: "Coffee Maker", Description: "Automatic drip coffee maker", Price: 89.99, Quantity: 30, CategoryID: 4},
		{Name: "Cookware Set", Description: "10-piece non-stick cookware set", Price: 199.99, Quantity: 25, CategoryID: 4},
		{Name: "Blender", Description: "High-speed countertop blender", Price: 79.99, Quantity: 40, CategoryID: 4},
		{Name: "Bedding Set", Description: "Queen size cotton bedding set", Price: 129.99, Quantity: 35, CategoryID: 4},
		{Name: "Smart Light Bulbs", Description: "WiFi-enabled color changing bulbs", Price: 49.99, Quantity: 60, CategoryID: 4},

		// Sports
		{Name: "Yoga Mat", Description: "Non-slip exercise yoga mat", Price: 29.99, Quantity: 100, CategoryID: 5},
		{Name: "Dumbbells", Description: "Pair of 5kg dumbbells", Price: 39.99, Quantity: 75, CategoryID: 5},
		{Name: "Basketball", Description: "Official size basketball", Price: 24.99, Quantity: 50, CategoryID: 5},
		{Name: "Tennis Racket", Description: "Professional tennis racket", Price: 149.99, Quantity: 30, CategoryID: 5},
		{Name: "Fitness Tracker", Description: "Activity and sleep tracking band", Price: 99.99, Quantity: 65, CategoryID: 5},
	}

	for _, product := range products {
		var existingProduct models.Product
		result := db.Where("name = ?", product.Name).First(&existingProduct)
		if result.RowsAffected == 0 {
			db.Create(&product)
			log.Printf("Created product: %s\n", product.Name)
		} else {
			log.Printf("Product already exists: %s\n", product.Name)
		}
	}

	log.Println("Finished seeding products successfully!")
}
