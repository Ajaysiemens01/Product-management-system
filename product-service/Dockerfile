# Use Golang image
FROM golang:1.23.4

# Set working directory
WORKDIR /app

# Copy source files
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o product-service main.go

# Expose service port
EXPOSE 8081

# Run the service
CMD ["./product-service"]
