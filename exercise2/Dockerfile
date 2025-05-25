FROM sbtscala/scala-sbt:eclipse-temurin-17.0.14_7_1.10.11_3.6.4 AS builder

WORKDIR /app

# Copy build files
COPY build.sbt /app/
COPY project /app/project/

# Copy only necessary source code files
COPY app/controllers/ /app/app/controllers/
COPY app/models/ /app/app/models/
COPY conf/application.conf /app/conf/application.conf
COPY conf/routes /app/conf/routes

# Build the application
RUN sbt stage

FROM eclipse-temurin:17-jre-jammy

WORKDIR /app

# Create a non-root user
RUN groupadd -r appuser && useradd -r -g appuser -m appuser

# Copy the built application
COPY --from=builder /app/target/universal/stage /app

# Copy the startup script
COPY startup.sh /app/
RUN chmod +x /app/startup.sh

# Set proper permissions
RUN chown -R appuser:appuser /app

# Expose the application port
EXPOSE 9000

# Switch to non-root user
USER appuser

CMD ["/app/startup.sh"]