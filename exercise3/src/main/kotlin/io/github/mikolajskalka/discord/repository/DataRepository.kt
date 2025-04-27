package io.github.mikolajskalka.discord.repository

import io.github.mikolajskalka.discord.models.Category
import io.github.mikolajskalka.discord.models.Product

/**
 * In-memory repository for product and category data
 */
object DataRepository {
    private val categories = listOf(
        Category(1, "Electronics", "Electronic devices and gadgets"),
        Category(2, "Clothing", "Fashion items and apparel"),
        Category(3, "Books", "Books and literature"),
        Category(4, "Home & Kitchen", "Home decor and kitchen items"),
        Category(5, "Sports & Outdoors", "Sports equipment and outdoor gear")
    )
    
    private val products = listOf(
        // Electronics
        Product(1, "Smartphone", "Latest model smartphone with advanced features", 799.99, 1),
        Product(2, "Laptop", "High-performance laptop for professionals", 1299.99, 1),
        Product(3, "Wireless Earbuds", "Premium wireless earbuds with noise cancellation", 149.99, 1),
        
        // Clothing
        Product(4, "T-Shirt", "Comfortable cotton t-shirt", 19.99, 2),
        Product(5, "Jeans", "Classic blue jeans", 49.99, 2),
        Product(6, "Sneakers", "Stylish and comfortable sneakers", 79.99, 2),
        
        // Books
        Product(7, "Programming Guide", "Comprehensive programming guide for beginners", 34.99, 3),
        Product(8, "Novel", "Bestselling fiction novel", 24.99, 3),
        Product(9, "Cookbook", "Collection of delicious recipes", 29.99, 3),
        
        // Home & Kitchen
        Product(10, "Coffee Maker", "Automatic coffee maker with timer", 89.99, 4),
        Product(11, "Blender", "High-speed blender for smoothies", 59.99, 4),
        Product(12, "Bedding Set", "Luxury bedding set with pillowcases", 99.99, 4),
        
        // Sports & Outdoors
        Product(13, "Yoga Mat", "Non-slip yoga mat for exercise", 29.99, 5),
        Product(14, "Dumbbell Set", "Adjustable dumbbell set for home workouts", 129.99, 5),
        Product(15, "Hiking Backpack", "Durable backpack for hiking and camping", 69.99, 5)
    )
    
    fun getAllCategories(): List<Category> = categories
    
    fun getCategoryById(id: Int): Category? = categories.find { it.id == id }
    
    fun getAllProducts(): List<Product> = products
    
    fun getProductById(id: Int): Product? = products.find { it.id == id }
    
    fun getProductsByCategoryId(categoryId: Int): List<Product> = products.filter { it.categoryId == categoryId }
}