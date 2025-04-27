package models

case class Product(id: Long, name: String, description: String, price: Double, categoryId: Long)

object ProductRepository {
  private var products = Vector(
    Product(1L, "iPhone", "Latest smartphone from Apple", 999.99, 1L),
    Product(2L, "MacBook Pro", "Powerful laptop for professionals", 1999.99, 1L),
    Product(3L, "Coffee Maker", "Premium coffee machine", 199.99, 2L),
    Product(4L, "Running Shoes", "Comfortable shoes for running", 89.99, 3L)
  )
  
  def findAll(): Seq[Product] = products
  
  def findById(id: Long): Option[Product] = products.find(_.id == id)
  
  def create(product: Product): Product = {
    val nextId = if (products.isEmpty) 1L else products.map(_.id).max + 1
    val newProduct = product.copy(id = nextId)
    products = products :+ newProduct
    newProduct
  }
  
  def update(product: Product): Option[Product] = {
    findById(product.id) match {
      case Some(_) =>
        products = products.filterNot(_.id == product.id) :+ product
        Some(product)
      case None => None
    }
  }
  
  def delete(id: Long): Boolean = {
    val initialSize = products.size
    products = products.filterNot(_.id == id)
    initialSize > products.size
  }
}