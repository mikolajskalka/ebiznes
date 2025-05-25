FROM gradle:8.6-jdk21 AS build

WORKDIR /app
# Copy only necessary files for the build
COPY build.gradle.kts settings.gradle.kts gradlew gradlew.bat ./
COPY gradle ./gradle
COPY src ./src
RUN gradle --no-daemon build

FROM openjdk:21-slim

WORKDIR /app

# Create a non-root user to run the application
RUN groupadd -r appuser && useradd -r -g appuser appuser

COPY --from=build /app/build/libs/*.jar app.jar

# Set proper permissions for the application files
RUN chown -R appuser:appuser /app

# Use environment variable from outside the container for security
ENV DISCORD_TOKEN="discord_token"

EXPOSE 8080

# Switch to non-root user
USER appuser

ENTRYPOINT ["java", "-jar", "app.jar"]