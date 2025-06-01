FROM golang:1.20-alpine AS builder
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o app .

# ----------- Run stage -----------
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/app .

# Copy config files
COPY manifest/config ./config

# Copy Swagger UI static files
COPY resource/public/swagger ./resource/public/swagger

EXPOSE 8000

# Start the application with config
CMD ["./app", "-c", "config/config.yaml"]