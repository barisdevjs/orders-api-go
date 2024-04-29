FROM golang:latest

WORKDIR /

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-orders-api-go

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["/docker-orders-api-go"]
