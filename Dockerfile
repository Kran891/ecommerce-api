# -------- Stage 1: Build --------
FROM golang:1.24-alpine AS builder

# Install git and tzdata
RUN apk update && apk add --no-cache git tzdata

# Set working directory
WORKDIR /app

# Cache Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build Go binary
RUN go build -o ecommerce-api ./cmd/main.go

# -------- Stage 2: Run --------
FROM alpine:latest

# Install timezone data
RUN apk --no-cache add ca-certificates tzdata

# Set timezone (optional)
ENV TZ=Asia/Kolkata

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/ecommerce-api .

# Set the entrypoint
CMD ["./ecommerce-api"]
