# Use a specific Golang version
FROM golang:1.22.0-alpine as builder

# Set working directory
WORKDIR /app

# Copy your source code
COPY . .

# Install dependencies
RUN go mod download

# Build the Go binary (replace "main.go" with your actual entry point)
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o EduNexus  cmd/EduNexus/main.go

# Use a smaller image for production
FROM alpine:latest

# Set working directory
WORKDIR /app


# Copy the .env file
COPY --from=builder /app/.env /app/.env

# Copy the binary
COPY --from=builder /app/EduNexus /app/EduNexus

# Expose the port
EXPOSE 8080 

# Start the EduNexus
CMD ["/app/EduNexus"]

