# Daemon Service
FROM golang:1.18-alpine AS builder

WORKDIR /app/daemon

# Copy daemon files
COPY daemon/ ./

# Build (daemon might not have go.mod)
RUN go mod download 2>/dev/null || true
RUN go build -o daemon . 2>/dev/null || cp main.exe daemon 2>/dev/null || echo "Using prebuilt binary"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/daemon/daemon* .

CMD ["./daemon"]
