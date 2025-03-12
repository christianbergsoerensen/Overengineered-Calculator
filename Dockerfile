# Stage 1: Build the Go application
FROM golang:1.24-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

# Copy the source code into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o main ./cmd/main.go

FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/

ENV PORT=8080
ENV DATABASE_URL ""
# Run the binary
CMD ["./main"]