# Use official Go image
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build binary
RUN go build -o todo

# Default command
ENTRYPOINT ["./todo"]
