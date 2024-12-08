FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o eco-api ./cmd/main.go

FROM ubuntu:noble-20241118.1

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && apt-get clean
WORKDIR /root/
COPY --from=builder /app/eco-api .
EXPOSE 5000
CMD ["./eco-api"]