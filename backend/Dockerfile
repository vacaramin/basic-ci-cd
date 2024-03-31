FROM golang:alpine AS builder

WORKDIR /app

# Copy go mod and sum files for installing dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .


# Install GCC
RUN apk add gcc musl-dev

# Build the application
RUN CGO_ENABLED=1 go build -o main .

EXPOSE 80

# Using lighter image for running the application build
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY ./.env .
COPY users.db .

CMD ["./main"]
