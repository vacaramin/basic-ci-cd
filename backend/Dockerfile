FROM golang:alpine AS builder

WORKDIR /app

# Copy go mod and sum files for installing dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

EXPOSE 8080

# Using lighter image for running the application build
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
