PROJECT_NAME ?= chi-go-template

.PHONY: init build build-web build-worker run web worker test vet fmt migrate-up migrate-down

init:
	@bash scripts/init.sh $(MODULE)

build:
	go build ./...

build-web:
	mkdir -p bin
	go build -o bin/$(PROJECT_NAME) ./cmd/web

build-worker:
	mkdir -p bin
	go build -o bin/$(PROJECT_NAME)-worker ./cmd/worker

web:
	go run ./cmd/web

worker:
	go run ./cmd/worker

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down