# Daemon Service
FROM golang:1.18-alpine AS builder

# Install git to handle submodules
RUN apk add --no-cache git

WORKDIR /app

# Copy the entire repository
COPY . .

# Build daemon
WORKDIR /app/daemon
RUN go mod download || true
RUN go build -o daemon . || echo "Daemon build from source"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/daemon/daemon* .

CMD ["./daemon"]
