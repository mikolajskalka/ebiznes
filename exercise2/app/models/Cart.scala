package models

case class CartItem(id: Long, productId: Long, quantity: Int)
case class Cart(id: Long, items: Seq[CartItem] = Seq.empty)

object CartRepository {
  private var carts = Vector(
    Cart(1L, Seq(CartItem(1L, 1L, 2), CartItem(2L, 3L, 1))),
    Cart(2L, Seq(CartItem(3L, 2L, 1)))
  )
  
  private var nextCartItemId: Long = 4L
  
  def findAll(): Seq[Cart] = carts
  
  def findById(id: Long): Option[Cart] = carts.find(_.id == id)
  
  def create(cart: Cart): Cart = {
    val nextId = if (carts.isEmpty) 1L else carts.map(_.id).max + 1
    val newCart = cart.copy(id = nextId)
    carts = carts :+ newCart
    newCart
  }
  
  def update(cart: Cart): Option[Cart] = {
    findById(cart.id) match {
      case Some(_) =>
        carts = carts.filterNot(_.id == cart.id) :+ cart
        Some(cart)
      case None => None
    }
  }
  
  def delete(id: Long): Boolean = {
    val initialSize = carts.size
    carts = carts.filterNot(_.id == id)
    initialSize > carts.size
  }
  
  // Additional methods for cart items
  def addItemToCart(cartId: Long, productId: Long, quantity: Int): Option[Cart] = {
    findById(cartId).map { cart =>
      val existingItemOpt = cart.items.find(_.productId == productId)
      
      val updatedItems = existingItemOpt match {
        case Some(existingItem) =>
          // Update quantity of existing item
          cart.items.map { item =>
            if (item.productId == productId) {
              item.copy(quantity = item.quantity + quantity)
            } else {
              item
            }
          }
        case None =>
          // Add new item
          val newItem = CartItem(nextCartItemId, productId, quantity)
          nextCartItemId += 1
          cart.items :+ newItem
      }
      
      val updatedCart = cart.copy(items = updatedItems)
      carts = carts.filterNot(_.id == cartId) :+ updatedCart
      updatedCart
    }
  }
  
  def removeItemFromCart(cartId: Long, itemId: Long): Option[Cart] = {
    findById(cartId).map { cart =>
      val updatedItems = cart.items.filterNot(_.id == itemId)
      val updatedCart = cart.copy(items = updatedItems)
      carts = carts.filterNot(_.id == cartId) :+ updatedCart
      updatedCart
    }
  }
}