PROJECT_DIR=$(shell pwd)

build:
	@go build -o bank_manager cmd/bank_manager/main.go

run: build
	@BANK_MANAGER=$(PROJECT_DIR) ./bank_manager

test:
	@BANK_MANAGER=$(PROJECT_DIR) go test ./...

fmt:
	@go fmt ./...
