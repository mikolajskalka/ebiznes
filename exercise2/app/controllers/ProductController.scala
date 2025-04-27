package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.{Product, ProductRepository}

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  
  // JSON format for Product
  implicit val productFormat: OFormat[Product] = Json.format[Product]
  
  // Get all products
  def index(): Action[AnyContent] = Action { implicit request =>
    val products = ProductRepository.findAll()
    Ok(Json.toJson(products))
  }
  
  // Get product by id
  def show(id: Long): Action[AnyContent] = Action { implicit request =>
    ProductRepository.findById(id) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("message" -> s"Product with id $id not found"))
    }
  }
  
  // Create new product
  def create(): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Product].fold(
      errors => {
        // Try to parse as a product without ID
        (request.body \ "name").asOpt[String].map { name =>
          val description = (request.body \ "description").asOpt[String].getOrElse("")
          val price = (request.body \ "price").asOpt[Double].getOrElse(0.0)
          val categoryId = (request.body \ "categoryId").asOpt[Long].getOrElse(0L)
          
          // Create with a temporary ID (will be replaced by the repository)
          val productToCreate = Product(0L, name, description, price, categoryId)
          
          val newProduct = ProductRepository.create(productToCreate)
          Created(Json.toJson(newProduct))
        }.getOrElse {
          BadRequest(Json.obj("message" -> JsError.toJson(errors)))
        }
      },
      product => {
        val newProduct = ProductRepository.create(product)
        Created(Json.toJson(newProduct))
      }
    )
  }
  
  // Update existing product
  def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Product].fold(
      errors => {
        // Try to parse as a product without ID
        (request.body \ "name").asOpt[String].map { name =>
          val description = (request.body \ "description").asOpt[String].getOrElse("")
          val price = (request.body \ "price").asOpt[Double].getOrElse(0.0)
          val categoryId = (request.body \ "categoryId").asOpt[Long].getOrElse(0L)
          
          val productToUpdate = Product(id, name, description, price, categoryId)
          
          ProductRepository.update(productToUpdate) match {
            case Some(updatedProduct) => Ok(Json.toJson(updatedProduct))
            case None => NotFound(Json.obj("message" -> s"Product with id $id not found"))
          }
        }.getOrElse {
          BadRequest(Json.obj("message" -> JsError.toJson(errors)))
        }
      },
      product => {
        ProductRepository.update(product.copy(id = id)) match {
          case Some(updatedProduct) => Ok(Json.toJson(updatedProduct))
          case None => NotFound(Json.obj("message" -> s"Product with id $id not found"))
        }
      }
    )
  }
  
  // Delete product
  def delete(id: Long): Action[AnyContent] = Action { implicit request =>
    if (ProductRepository.delete(id)) {
      NoContent
    } else {
      NotFound(Json.obj("message" -> s"Product with id $id not found"))
    }
  }
}