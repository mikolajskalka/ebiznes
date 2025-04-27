package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.{Cart, CartItem, CartRepository}

@Singleton
class CartController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  
  // JSON format for CartItem and Cart
  implicit val cartItemFormat: OFormat[CartItem] = Json.format[CartItem]
  implicit val cartFormat: OFormat[Cart] = Json.format[Cart]
  
  // Get all carts
  def index(): Action[AnyContent] = Action { implicit request =>
    val carts = CartRepository.findAll()
    Ok(Json.toJson(carts))
  }
  
  // Get cart by id
  def show(id: Long): Action[AnyContent] = Action { implicit request =>
    CartRepository.findById(id) match {
      case Some(cart) => Ok(Json.toJson(cart))
      case None => NotFound(Json.obj("message" -> s"Cart with id $id not found"))
    }
  }
  
  // Create new cart
  def create(): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Cart].fold(
      errors => {
        // Try to parse as a cart without ID
        // For a new cart, we can accept an empty request or just the items array
        val items = (request.body \ "items").asOpt[JsArray]
          .map(_.value.map(_.as[CartItem]).toList)
          .getOrElse(List.empty[CartItem])
        
        // Create with a temporary ID (will be replaced by the repository)
        val cartToCreate = Cart(0L, items)
        
        val newCart = CartRepository.create(cartToCreate)
        Created(Json.toJson(newCart))
      },
      cart => {
        val newCart = CartRepository.create(cart)
        Created(Json.toJson(newCart))
      }
    )
  }
  
  // Update existing cart
  def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Cart].fold(
      errors => {
        // Try to parse as a cart without ID
        (request.body \ "items").asOpt[JsArray].map { itemsJson =>
          try {
            val items = itemsJson.value.map(_.as[CartItem]).toList
            val cartToUpdate = Cart(id, items)
            
            CartRepository.update(cartToUpdate) match {
              case Some(updatedCart) => Ok(Json.toJson(updatedCart))
              case None => NotFound(Json.obj("message" -> s"Cart with id $id not found"))
            }
          } catch {
            case e: JsResultException => 
              BadRequest(Json.obj("message" -> "Invalid items format in cart"))
          }
        }.getOrElse {
          BadRequest(Json.obj("message" -> JsError.toJson(errors)))
        }
      },
      cart => {
        CartRepository.update(cart.copy(id = id)) match {
          case Some(updatedCart) => Ok(Json.toJson(updatedCart))
          case None => NotFound(Json.obj("message" -> s"Cart with id $id not found"))
        }
      }
    )
  }
  
  // Delete cart
  def delete(id: Long): Action[AnyContent] = Action { implicit request =>
    if (CartRepository.delete(id)) {
      NoContent
    } else {
      NotFound(Json.obj("message" -> s"Cart with id $id not found"))
    }
  }
  
  // Remove item from cart
  def removeItem(cartId: Long, itemId: Long): Action[AnyContent] = Action { implicit request =>
    CartRepository.removeItemFromCart(cartId, itemId) match {
      case Some(updatedCart) => Ok(Json.toJson(updatedCart))
      case None => NotFound(Json.obj("message" -> s"Cart with id $cartId or item with id $itemId not found"))
    }
  }

  // Add item to cart
  def addItem(cartId: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    val productIdResult = (request.body \ "productId").validate[Long]
    val quantityResult = (request.body \ "quantity").validate[Int]
    
    (productIdResult, quantityResult) match {
      case (JsSuccess(productId, _), JsSuccess(quantity, _)) =>
        CartRepository.addItemToCart(cartId, productId, quantity) match {
          case Some(updatedCart) => Ok(Json.toJson(updatedCart))
          case None => NotFound(Json.obj("message" -> s"Cart with id $cartId not found"))
        }
      case _ =>
        BadRequest(Json.obj("message" -> "Invalid request body. Expected 'productId' and 'quantity' fields"))
    }
  }
}