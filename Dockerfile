FROM golang:1.21.4

# Create the /app/downloads directory
RUN mkdir -p /app/downloads

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the application code
COPY . .

# Build the application
RUN go build -o ./bin ./cmd

# Define the command to run the binary
CMD ["./bin"]
