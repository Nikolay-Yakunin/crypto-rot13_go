FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/crypto-rot13 ./cmd/crypto-rot13
COPY ./internal/rot13 ./internal/rot13
COPY ./static ./static
COPY .env .env

RUN go build -o crypto-rot13 ./cmd/crypto-rot13

CMD ["./crypto-rot13"]