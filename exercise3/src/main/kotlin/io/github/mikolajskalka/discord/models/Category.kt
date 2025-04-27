package io.github.mikolajskalka.discord.models

import kotlinx.serialization.Serializable

@Serializable
data class Category(
    val id: Int,
    val name: String,
    val description: String
)