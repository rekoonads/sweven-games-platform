# Multi-stage build for Sweven Games Backend
FROM golang:1.18-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy webrtc-proxy
COPY webrtc-proxy/ ./webrtc-proxy/
WORKDIR /app/webrtc-proxy

# Download dependencies and build
RUN go mod download
RUN go build -o /app/bin/webrtc-proxy .

# Copy daemon
WORKDIR /app/daemon
COPY daemon/ .
RUN go mod download || true
RUN go build -o /app/bin/daemon . || cp main.exe /app/bin/daemon

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy built binaries
COPY --from=builder /app/bin/ ./bin/

# Copy configuration files
COPY script/ ./script/
COPY package/ ./package/

# Expose ports
EXPOSE 8080 8088 8000

# Default to running webrtc-proxy
CMD ["./bin/webrtc-proxy"]
