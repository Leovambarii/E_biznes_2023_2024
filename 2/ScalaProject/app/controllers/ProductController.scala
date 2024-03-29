package controllers

import javax.inject._
import models._
import play.api.mvc._
import scala.collection.mutable
import play.api.libs.json._

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  private val products = new mutable.ListBuffer[Product]()
  products += Product(1, "Product 1", "First product", 1, 50)
  products += Product(2, "Product 2", "Second product", 2, 45)
  products += Product(3, "Product 3", "Third product", 2, 115)
  products += Product(4, "Product 4", "Test product", 2, 45)
  products += Product(5, "Product 5", "Final product", 3, 155)

  def showAll(): Action[AnyContent] = Action {
    if (products.isEmpty) {
      NoContent
    }
    else {
      Ok(Json.toJson(products))
    }
  }

  def showById(productId: Long): Action[AnyContent] = Action {
    products.find(_.id == productId) match {
      case Some(product) => Ok(Json.toJson(product))
      case None => NotFound(Json.obj("message" -> "Product not found"))
    }
  }

  def updateItem(productId: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[JsObject].fold(
      errors => {
        BadRequest(Json.obj("message" -> JsError.toJson(errors)))
      },
      productFields => {
        products.indexWhere(_.id == productId) match {
          case -1 => NotFound(Json.obj("message" -> "Product not found"))
          case index =>
            val currProduct = products(index)
            val updatedProduct = currProduct.copy(
              name = (productFields \ "name").asOpt[String].getOrElse(currProduct.name),
              description = (productFields \ "description").asOpt[String].getOrElse(currProduct.description),
              category = (productFields \ "category").asOpt[Int].getOrElse(currProduct.category),
              price = (productFields \ "price").asOpt[Int].getOrElse(currProduct.price)
            )
            products.update(index, updatedProduct)
            Ok(Json.obj("message" -> "Product updated successfully", "product" -> updatedProduct))
        }
      }
    )
  }

  def deleteItem(productId: Long): Action[AnyContent] = Action {
    products.indexWhere(_.id == productId) match {
      case -1 => NotFound(Json.obj("message" -> "Product not found"))
      case index =>
        products.remove(index)
        Ok(Json.obj("message" -> "Product deleted successfully"))
    }
  }

  def addItem(): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[NewProduct].fold(
      errors => BadRequest(Json.obj("message" -> JsError.toJson(errors))),
      product => {
        val newId = if (products.isEmpty) 1 else products.map(_.id).max + 1
        val newProduct = Product(newId, product.name, product.description, product.category, product.price)
        products += newProduct
        Created(Json.obj("message" -> "Product added successfully", "product" -> newProduct))
      }
    )
  }
}
