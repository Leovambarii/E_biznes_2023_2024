package com.ktor_project

import dev.kord.common.entity.Snowflake
import dev.kord.core.Kord
import dev.kord.core.entity.channel.TextChannel
import dev.kord.core.event.message.MessageCreateEvent
import dev.kord.core.on
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

data class Category(val id: Int, var name: String)
data class Product(val id: Int, var name: String, var description: String, var category: Int, var price: Int)

suspend fun main() {
    val token = System.getenv("DISCORD_BOT_TOKEN") ?: error("Token not found in environment variables...\nPlease create it before running the project...")
    val kord = Kord(token)
    val channelId = Snowflake("1225585240035168348")

    val categories = listOf(
        Category(1, "A"),
        Category(2, "B"),
        Category(3, "C")
    )

    val products = listOf(
        Product(1, "Product 1", "Description 1", 1, 55),
        Product(2, "Product 2", "Description 2", 1, 18),
        Product(3, "Product 3", "Description 3", 2, 12),
        Product(4, "Product 4", "Description 4", 3, 1)
    )

    kord.on<MessageCreateEvent> {
        if (message.author?.isBot != false) return@on
        if (message.channelId != channelId) return@on
        if (!message.content.startsWith("!")) return@on

        println("Message received: ${message.content}")
        when (message.content.split(" ")[0]) {
            "!products" -> {
                val formatted = products.joinToString("\n") {
                    product -> String.format("%d   '%s'   '%s'   %d   %d", product.id, product.name, product.description, product.category, product.price)
                }
                message.channel.createMessage("Available products (id, name, description, category, price):\n${formatted}")
            }
            "!categories" -> {
                val formatted = categories.joinToString(separator = "\n") {
                    category -> String.format("%d   '%s'", category.id, category.name)
                }
                message.channel.createMessage("Available categories (id, name):\n${formatted}")
            }
            "!category" -> {
                if (message.content.split(" ").size != 2) {
                    message.channel.createMessage("Invalid command...\n\nUse !category <id>")
                    return@on
                }

                val id = message.content.split(" ")[1].toIntOrNull()
                if (id == null || categories.none { it.id == id }) {
                    message.channel.createMessage("Category id not found...")
                    return@on
                }

                val filteredProducts = products.filter { it.category == id }
                if (filteredProducts.isEmpty())
                    message.channel.createMessage("No products found for category id=${id}...")
                else {
                    val formatted = filteredProducts.joinToString("\n") {
                            product -> String.format("%d   '%s'   '%s'   %d   %d", product.id, product.name, product.description, product.category, product.price)
                    }
                    message.channel.createMessage("Products for category id=${id} (id, name, description, category, price):\n${formatted}")
                }
            } else -> message.channel.createMessage(
                "Unknown command...\n\n" +
                        "Available commands:\n" +
                        "!products           - List all available products\n" +
                        "!categories        - List available categories\n" +
                        "!category <id>  - List products with given category id"
            )
        }
    }

    embeddedServer(Netty, port = 8080, host = "0.0.0.0") {
        sendMessage(kord, channelId)
    }.start(wait = false)

    kord.login {
        @OptIn(PrivilegedIntent::class)
        intents += Intent.MessageContent
    }
}

fun Application.sendMessage(kord: Kord, channelId: Snowflake) {
    routing {
        post("/send") {
            val message = call.receive<String>()
            val channel: TextChannel? = kord.getChannelOf<TextChannel>(channelId)

            channel?.createMessage(message)
            call.respond(HttpStatusCode.OK, "Message sent successfully")
        }
    }
}