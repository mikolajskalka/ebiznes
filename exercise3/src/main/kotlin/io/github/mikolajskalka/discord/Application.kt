package io.github.mikolajskalka.discord

import io.github.mikolajskalka.discord.bot.DiscordBot
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

/**
 * Main application class that starts the Ktor server and Discord bot
 */
fun main(args: Array<String>) {
    // Get Discord token from environment variable or use a placeholder
    val discordToken = System.getenv("DISCORD_TOKEN") ?: "discord_token"
    
    // Create Discord bot instance
    val discordBot = DiscordBot(discordToken)
    
    try {
        // Start the Discord bot
        discordBot.start()
        
        // Start the Ktor server
        embeddedServer(Netty, port = 8080) {
            configureRouting(discordBot)
        }.start(wait = true)
    } catch (e: Exception) {
        println("Error starting application: ${e.message}")
    } finally {
        // Shutdown resources when application stops
        discordBot.stop()
    }
}

/**
 * Configure Ktor routing
 */
fun Application.configureRouting(discordBot: DiscordBot) {
    routing {
        // Health check endpoint
        get("/health") {
            call.respondText("Bot is running!")
        }
        
        // Endpoint to send a test message to a Discord channel
        get("/send-discord-message") {
            val channelId = call.request.queryParameters["channelId"] ?: run {
                call.respondText("Missing channelId parameter")
                return@get
            }
            
            val message = call.request.queryParameters["message"] ?: run {
                call.respondText("Missing message parameter")
                return@get
            }
            
            discordBot.sendMessage(channelId, message)
            call.respondText("Message sent to Discord channel $channelId")
        }
    }
}