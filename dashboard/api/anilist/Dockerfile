# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

# Install git if you fetch modules from VCS
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o app .

# Stage 2: Minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/app .

# Expose port (same as in your Go server)
EXPOSE 3000

# Run the binary
CMD ["./app"]
