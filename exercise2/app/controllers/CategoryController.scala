package controllers

import javax.inject._
import play.api.mvc._
import play.api.libs.json._
import models.{Category, CategoryRepository}

@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  
  // JSON format for Category
  implicit val categoryFormat: OFormat[Category] = Json.format[Category]
  
  // Get all categories
  def index(): Action[AnyContent] = Action { implicit request =>
    val categories = CategoryRepository.findAll()
    Ok(Json.toJson(categories))
  }
  
  // Get category by id
  def show(id: Long): Action[AnyContent] = Action { implicit request =>
    CategoryRepository.findById(id) match {
      case Some(category) => Ok(Json.toJson(category))
      case None => NotFound(Json.obj("message" -> s"Category with id $id not found"))
    }
  }
  
  // Create new category
  def create(): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Category].fold(
      errors => {
        // Try to parse as a category without ID
        (request.body \ "name").asOpt[String].map { name =>
          val description = (request.body \ "description").asOpt[String].getOrElse("")
          
          // Create with a temporary ID (will be replaced by the repository)
          val categoryToCreate = Category(0L, name, description)
          
          val newCategory = CategoryRepository.create(categoryToCreate)
          Created(Json.toJson(newCategory))
        }.getOrElse {
          BadRequest(Json.obj("message" -> JsError.toJson(errors)))
        }
      },
      category => {
        val newCategory = CategoryRepository.create(category)
        Created(Json.toJson(newCategory))
      }
    )
  }
  
  // Update existing category
  def update(id: Long): Action[JsValue] = Action(parse.json) { implicit request =>
    request.body.validate[Category].fold(
      errors => {
        // Try to parse as a category without ID
        (request.body \ "name").asOpt[String].map { name =>
          val description = (request.body \ "description").asOpt[String].getOrElse("")
          
          val categoryToUpdate = Category(id, name, description)
          
          CategoryRepository.update(categoryToUpdate) match {
            case Some(updatedCategory) => Ok(Json.toJson(updatedCategory))
            case None => NotFound(Json.obj("message" -> s"Category with id $id not found"))
          }
        }.getOrElse {
          BadRequest(Json.obj("message" -> JsError.toJson(errors)))
        }
      },
      category => {
        CategoryRepository.update(category.copy(id = id)) match {
          case Some(updatedCategory) => Ok(Json.toJson(updatedCategory))
          case None => NotFound(Json.obj("message" -> s"Category with id $id not found"))
        }
      }
    )
  }
  
  // Delete category
  def delete(id: Long): Action[AnyContent] = Action { implicit request =>
    if (CategoryRepository.delete(id)) {
      NoContent
    } else {
      NotFound(Json.obj("message" -> s"Category with id $id not found"))
    }
  }
}