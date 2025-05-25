FROM golang:1.24-alpine AS build

WORKDIR /app

# Install necessary dependencies for CGO and SQLite
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy only necessary source code files
COPY main.go ./
COPY controllers/ ./controllers/
COPY database/ ./database/
COPY models/ ./models/
COPY routes/ ./routes/
COPY scripts/ ./scripts/

# Build the Go application and seed script - CGO is enabled for SQLite
RUN CGO_ENABLED=1 go build -o main . && \
    CGO_ENABLED=1 go build -o seed_db ./scripts/seed_products.go

# Create a minimal image for running the application
FROM alpine:3.21.3

WORKDIR /app

# Install runtime dependencies for SQLite
RUN apk --no-cache add ca-certificates sqlite

# Create a non-root user to run the application
RUN addgroup -S appgroup && adduser -S -G appgroup appuser

# Copy the binaries from the build stage
COPY --from=build /app/main .
COPY --from=build /app/seed_db .

# Create a database directory, set permissions, and seed the database
RUN mkdir -p /app/data && \
    touch /app/data/shop.db && \
    chown -R appuser:appgroup /app && \
    chmod +x /app/main && \
    chmod +x /app/seed_db && \
    chmod 644 /app/data/shop.db && \
    chmod 755 /app/data && \
    ./seed_db && \
    chown appuser:appgroup /app/data/shop.db

# Expose port 8080
EXPOSE 8080

# Create a data directory for the SQLite database
VOLUME /app/data

# Switch to non-root user
USER appuser

# Command to run the application
CMD ["./main"]