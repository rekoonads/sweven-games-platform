# Signaling Server Service
FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY signaling-server/ ./signaling-server/
WORKDIR /app/signaling-server

RUN go mod download
RUN go build -o signaling .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/signaling-server/signaling .

EXPOSE 8088 8000

CMD ["./signaling", "--websocket", "8088", "--grpc", "8000"]
