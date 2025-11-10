# WebRTC Proxy Service
FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY webrtc-proxy/ ./webrtc-proxy/
WORKDIR /app/webrtc-proxy

RUN go mod download
RUN go build -o webrtc-proxy .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/webrtc-proxy/webrtc-proxy .

EXPOSE 8080

CMD ["./webrtc-proxy"]
