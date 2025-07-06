FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/crypto-rot13 ./cmd/crypto-rot13
COPY ./internal/crypto ./internal/crypto
COPY ./static ./static
COPY ./pkg/middleware ./pkg/middleware
COPY .env .env

RUN go build -o crypto-rot13 ./cmd/crypto-rot13

CMD ["./crypto-rot13"]