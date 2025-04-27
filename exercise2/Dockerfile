FROM sbtscala/scala-sbt:eclipse-temurin-17.0.14_7_1.10.11_3.6.4 AS builder

WORKDIR /app

# Copy build files
COPY build.sbt /app/
COPY project /app/project/

# Copy source code
COPY app /app/app/
COPY conf /app/conf/

# Build the application
RUN sbt stage

FROM eclipse-temurin:17-jre-jammy

WORKDIR /app

# Copy the built application
COPY --from=builder /app/target/universal/stage /app

# Copy the startup script
COPY startup.sh /app/
RUN chmod +x /app/startup.sh

# Expose the application port
EXPOSE 9000

CMD ["/app/startup.sh"]