# WebRTC Proxy Service
FROM golang:1.18-alpine AS builder

WORKDIR /app/webrtc-proxy

# Copy webrtc-proxy files
COPY webrtc-proxy/go.mod webrtc-proxy/go.sum ./
RUN go mod download

# Copy the rest of the source
COPY webrtc-proxy/ ./

# Build
RUN go build -o webrtc-proxy .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/webrtc-proxy/webrtc-proxy .

EXPOSE 8080

CMD ["./webrtc-proxy"]
