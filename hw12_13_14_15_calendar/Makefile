BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./configs/config.yaml

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
update-generate-tools:
	go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
generate:
	protoc -I . \
			--go_out ./api/calendar \
			--go-grpc_out ./api/calendar \
			--grpc-gateway_out ./api/calendar \
      		--grpc-gateway_opt logtostderr=true \
        	--grpc-gateway_opt generate_unbound_methods=true \
			./proto/*.proto


.PHONY: build run build-img run-img version test lint
