package models

import play.api.libs.json.Json

case class Product(id: Long, name: String, description: String, category: Int, price: Int)
object Product {
  implicit val productFormat = Json.format[Product]
}

case class NewProduct(name: String, description: String, category: Int, price: Int)
object NewProduct {
  implicit val productFormat = Json.format[NewProduct]
}
