FROM golang:1.24-alpine AS build

WORKDIR /app

# Install necessary dependencies for CGO and SQLite
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go application - CGO is enabled for SQLite
RUN CGO_ENABLED=1 go build -o main .

# Create a minimal image for running the application
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies for SQLite
RUN apk --no-cache add ca-certificates sqlite

# Copy the binary from the build stage
COPY --from=build /app/main .

# Copy SQLite database file if it exists
COPY --from=build /app/shop.db ./shop.db

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]