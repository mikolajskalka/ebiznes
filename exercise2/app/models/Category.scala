package models

case class Category(id: Long, name: String, description: String)

object CategoryRepository {
  private var categories = Vector(
    Category(1L, "Electronics", "Electronic devices and gadgets"),
    Category(2L, "Home Appliances", "Appliances for your home"),
    Category(3L, "Sports", "Sports equipment and clothing")
  )
  
  def findAll(): Seq[Category] = categories
  
  def findById(id: Long): Option[Category] = categories.find(_.id == id)
  
  def create(category: Category): Category = {
    val nextId = if (categories.isEmpty) 1L else categories.map(_.id).max + 1
    val newCategory = category.copy(id = nextId)
    categories = categories :+ newCategory
    newCategory
  }
  
  def update(category: Category): Option[Category] = {
    findById(category.id) match {
      case Some(_) =>
        categories = categories.filterNot(_.id == category.id) :+ category
        Some(category)
      case None => None
    }
  }
  
  def delete(id: Long): Boolean = {
    val initialSize = categories.size
    categories = categories.filterNot(_.id == id)
    initialSize > categories.size
  }
}