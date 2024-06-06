# Use the official Golang image for the build process
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a minimal base image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/app .

COPY .env .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]