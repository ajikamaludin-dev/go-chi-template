.PHONY: init build run web worker test vet fmt

init:
	@bash scripts/init.sh $(MODULE)

build:
	go build ./...

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