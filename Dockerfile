FROM golang:1.22-alpine AS builder

# Install git
RUN apk update && apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o tigerbeetle_api .

# Use a minimal base image for the final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/tigerbeetle_api .

# Expose port (if your application listens on a port)
# EXPOSE 8080

# Set the entrypoint
ENTRYPOINT ["./tigerbeetle_api"]
