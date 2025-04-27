package io.github.mikolajskalka.discord.bot

import io.github.mikolajskalka.discord.repository.DataRepository
import net.dv8tion.jda.api.EmbedBuilder
import net.dv8tion.jda.api.JDA
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.entities.MessageEmbed
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter
import net.dv8tion.jda.api.requests.GatewayIntent
import java.awt.Color

/**
 * Discord bot handler that implements the required features:
 * - Sends messages to Discord
 * - Receives messages from users
 * - Returns categories list on request
 * - Returns products by category on request
 */
class DiscordBot(private val token: String) : ListenerAdapter() {
    private lateinit var jda: JDA
    
    fun start() {
        jda = JDABuilder.createDefault(token)
            .enableIntents(GatewayIntent.MESSAGE_CONTENT)
            .addEventListeners(this)
            .build()
            .awaitReady()
        
        println("Discord bot started: ${jda.selfUser.name}")
    }
    
    fun stop() {
        if (::jda.isInitialized) {
            jda.shutdown()
            println("Discord bot stopped")
        }
    }
    
    override fun onMessageReceived(event: MessageReceivedEvent) {
        // Ignore messages from bots (including self)
        if (event.author.isBot) return
        
        val content = event.message.contentRaw
        val channel = event.channel
        
        when {
            // Command to list all categories
            content.startsWith("!categories") -> {
                val categories = DataRepository.getAllCategories()
                val embed = EmbedBuilder()
                    .setTitle("Available Categories")
                    .setColor(Color.GREEN)
                    .setDescription("Here are all available product categories")
                
                categories.forEach { category ->
                    embed.addField(
                        "${category.id}. ${category.name}",
                        category.description,
                        false
                    )
                }
                
                embed.setFooter("Use !products <category_id> to view products")
                channel.sendMessageEmbeds(embed.build()).queue()
            }
            
            // Command to list products by category ID
            content.startsWith("!products") -> {
                val args = content.split(" ")
                if (args.size < 2) {
                    channel.sendMessage("Please specify a category ID. Example: !products 1").queue()
                    return
                }
                
                val categoryId = args[1].toIntOrNull()
                if (categoryId == null) {
                    channel.sendMessage("Please enter a valid category ID (number)").queue()
                    return
                }
                
                val category = DataRepository.getCategoryById(categoryId)
                if (category == null) {
                    channel.sendMessage("Category with ID $categoryId not found").queue()
                    return
                }
                
                val products = DataRepository.getProductsByCategoryId(categoryId)
                if (products.isEmpty()) {
                    channel.sendMessage("No products found in category ${category.name}").queue()
                    return
                }
                
                val embed = EmbedBuilder()
                    .setTitle("Products in ${category.name}")
                    .setColor(Color.BLUE)
                    .setDescription("Here are all products in ${category.name}")
                
                products.forEach { product ->
                    embed.addField(
                        "${product.id}. ${product.name}",
                        "${product.description}\nPrice: $${product.price}",
                        false
                    )
                }
                
                channel.sendMessageEmbeds(embed.build()).queue()
            }
            
            // Command to show help
            content.startsWith("!help") -> {
                val helpEmbed = EmbedBuilder()
                    .setTitle("Bot Commands")
                    .setColor(Color.YELLOW)
                    .setDescription("Here are the available commands:")
                    .addField("!categories", "Lists all product categories", false)
                    .addField("!products <category_id>", "Lists all products in a specific category", false)
                    .addField("!help", "Displays this help message", false)
                    .build()
                
                channel.sendMessageEmbeds(helpEmbed).queue()
            }
        }
    }
    
    fun sendMessage(channelId: String, message: String) {
        jda.getTextChannelById(channelId)?.sendMessage(message)?.queue()
    }
    
    fun sendEmbed(channelId: String, embed: MessageEmbed) {
        jda.getTextChannelById(channelId)?.sendMessageEmbeds(embed)?.queue()
    }
}