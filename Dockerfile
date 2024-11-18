# Use the official Golang image to build the Go app
FROM golang:1.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app (this will be the client that runs the REST API)
RUN CGO_ENABLED=0 GOOS=linux go build -o client .

# Create a small image with only the built binary
FROM alpine:latest  

# Install necessary libraries for running the Go binary
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/client .

# Expose the port the client will run on (adjust accordingly)
EXPOSE 8080

# Command to run the client
CMD ["./client"]
