// Define versions as extra properties
val kotlinVersion by extra("1.9.21")
val ktorVersion by extra("2.3.7")
val logbackVersion by extra("1.4.11")
val jdaVersion by extra("5.0.0-beta.18")
val serializationVersion by extra("1.6.0") 
val coroutinesVersion by extra("1.7.3")

plugins {
    kotlin("jvm") version "1.9.21" 
    kotlin("plugin.serialization") version "1.9.21"
    application
}

group = "io.github.mikolajskalka"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    // Kotlin standard library
    implementation("org.jetbrains.kotlin:kotlin-stdlib:$kotlinVersion")
    
    // Ktor server dependencies
    implementation("io.ktor:ktor-server-core:$ktorVersion")
    implementation("io.ktor:ktor-server-netty:$ktorVersion")
    
    // Ktor client dependencies
    implementation("io.ktor:ktor-client-core:$ktorVersion")
    implementation("io.ktor:ktor-client-cio:$ktorVersion")
    implementation("io.ktor:ktor-client-content-negotiation:$ktorVersion")
    implementation("io.ktor:ktor-serialization-kotlinx-json:$ktorVersion")
    
    // Logging
    implementation("ch.qos.logback:logback-classic:$logbackVersion")
    
    // Discord API (JDA)
    implementation("net.dv8tion:JDA:$jdaVersion")
    
    // Serialization
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:$serializationVersion")
    
    // Coroutines
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:$coroutinesVersion")

    // Tests
    testImplementation("io.ktor:ktor-server-test-host:$ktorVersion")
    testImplementation("org.jetbrains.kotlin:kotlin-test:$kotlinVersion")
}

application {
    mainClass.set("io.github.mikolajskalka.discord.ApplicationKt")
}

tasks.jar {
    manifest {
        attributes["Main-Class"] = "io.github.mikolajskalka.discord.ApplicationKt"
    }
    
    // Include all the dependencies in the jar
    from(configurations.runtimeClasspath.get().map { if (it.isDirectory) it else zipTree(it) })
    duplicatesStrategy = DuplicatesStrategy.EXCLUDE
}