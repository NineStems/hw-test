BIN := "./bin/calendar"
BIN_SCHEDULER= "./bin/calendar_scheduler"
BIN_SENDER= "./bin/calendar_sender"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/scheduler
	go build -v -o $(BIN_SENDER) -ldflags "$(LDFLAGS)" ./cmd/sender

run: build
	$(BIN) -config ./configs/calendar_config.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

goose-get:
	go install github.com/pressly/goose/v3/cmd/goose@latest

goose-up:
	goose -dir="migrations" postgres "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable" up

goose-down:
	goose -dir="migrations" postgres "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable" down

goose-version:
	goose postgres "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable" version

generate:
	protoc -I . \
			--go_out . \
			--go-grpc_out . \
			--grpc-gateway_out . \
      		--grpc-gateway_opt logtostderr=true \
        	--grpc-gateway_opt generate_unbound_methods=true \
			proto/api/v1/*.proto

docker-service-up:
	docker-compose -f docker-compose.yaml up -d --force-recreate

docker-service-down:
	docker-compose -f docker-compose.yaml down

.PHONY: build run build-img run-img version test lint

up: docker-service-up goose-get goose-up

down: goose-get goose-down docker-service-down