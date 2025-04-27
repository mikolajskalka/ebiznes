FROM gradle:8.6-jdk21 AS build

WORKDIR /app
COPY . .
RUN gradle --no-daemon build

FROM openjdk:21-slim

WORKDIR /app
COPY --from=build /app/build/libs/*.jar app.jar

ENV DISCORD_TOKEN="discord_token"

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "app.jar"]