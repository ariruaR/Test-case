# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/components /root/components
COPY --from=builder /app/docs /root/docs
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./app"]