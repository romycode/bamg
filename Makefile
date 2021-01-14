build:
	@go build -o bank_manager cmd/bank_manager/main.go

run: build
	@./bank_manager

test:
	@go test ./...

fmt:
	@go fmt ./...
