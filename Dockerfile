FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Install dependencies including curl
RUN apk update && apk add --no-cache git curl

# Download and install migrate tool to /usr/local/bin
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Create final minimal image
FROM alpine:3.19

WORKDIR /app

# Install postgresql-client for psql (optional)
RUN apk update && apk add --no-cache postgresql-client

# Copy migrate tool from builder stage
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

# Copy built binary from builder stage
COPY --from=builder /app/main .

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 3000

# Command to run the application
CMD ["./main"]