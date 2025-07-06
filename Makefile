CC=go
LD_FLAGS="-X main.Version=1.0.0 -X main.Commit=$(shell git rev-parse --short HEAD) -X main.BuildTime=$(shell date +%FT%T)"

BIN=crypto-rot13
CMD=cmd/crypto-rot13/main.go

PORT ?= $(shell [ -f .env ] && grep -E '^PORT=' .env | cut -d '=' -f2 || echo 8080)

all: build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(CC) build -ldflags $(LD_FLAGS) -o bin/$(BIN) $(CMD)

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(CC) build -o bin/$(BIN)-linux $(CMD)

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(CC) build -o bin/$(BIN).exe $(CMD)

build-macos:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(CC) build -o bin/$(BIN)-macos $(CMD)

run: build
	./bin/$(BIN)

clean:
	rm -f ./bin/*

style:
	find . -name "*.go" | xargs gofmt -s -w 
	find . -name "*.go" | xargs goimports -w 

test:
	$(CC) test ./... -v

lint:
	$(CC) vet ./...

docker:
	docker build -t $(BIN) .

docker-run:
	docker run --rm -it -p ${PORT}:${PORT} --env-file .env $(BIN)

docker-compose:
	docker-compose up --build

# coverage:
# 	$(CC) test ./... -coverprofile=coverage.out

# coverage-html: coverage
# 	$(CC) tool cover -html=coverage.out -o coverage.html

.PHONY: all build run clean test lint style docker docker-run docker-compose