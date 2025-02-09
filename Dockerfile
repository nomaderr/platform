# ======== STAGE 1: Build the Go Binary ========
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to take advantage of Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire backend folder
COPY backend /app/backend
COPY frontend /app/frontend

# Copy frontend main.go (as it acts as an entry point)
COPY frontend/main.go /app/main.go

# Set the correct module path
WORKDIR /app

# Build the Go application
RUN go build -o server main.go

# ======== STAGE 2: Create Final Minimal Image ========
FROM alpine:latest

# Set working directory inside the final container
WORKDIR /app

# Copy the compiled Go binary
COPY --from=builder /app/server .

# Copy frontend static & templates
COPY frontend/static /app/static
COPY frontend/templates /app/templates

# Expose the application port
EXPOSE 8080

# Start the Go application
CMD ["./server"]
