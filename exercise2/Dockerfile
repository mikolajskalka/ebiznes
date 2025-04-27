FROM sbtscala/scala-sbt:eclipse-temurin-17.0.14_7_1.10.11_3.6.4 AS builder

WORKDIR /app

# Copy build files
COPY build.sbt /app/
COPY project /app/project/

# Copy source code
COPY app /app/app/
COPY conf /app/conf/
# COPY public /app/public/

# Build the application
RUN sbt stage

# Create production image
FROM eclipse-temurin:17-jre-jammy

WORKDIR /app

# Copy the built application
COPY --from=builder /app/target/universal/stage /app

# Expose the application port
EXPOSE 9000

# Run the application
CMD ["/app/bin/e-shop-api", "-Dplay.http.secret.key=ad31779d4ee49d5ad5162bf1429c32e2e9933f3b"]