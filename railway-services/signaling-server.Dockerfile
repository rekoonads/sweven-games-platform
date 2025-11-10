# Signaling Server Service
FROM golang:1.18-alpine AS builder

WORKDIR /app/signaling-server

# Copy signaling server files directly
COPY signaling-server/go.mod signaling-server/go.sum ./
RUN go mod download

# Copy the rest of the signaling server source
COPY signaling-server/ ./

# Build from cmd directory
RUN go build -o signaling ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/signaling-server/signaling .

EXPOSE 8088 8000

CMD ["./signaling", "--websocket", "8088", "--grpc", "8000"]
