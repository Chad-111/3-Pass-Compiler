# Use the official Golang image as the base
FROM golang:1.23.0-alpine

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o service .

# Expose the port for each service
EXPOSE 8082

# Run the service
CMD ["./service"]
