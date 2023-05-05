.PHONY: run
run: cmd/spider/main.go
	go run cmd/spider/main.go

.PHONY: build
build: cmd/spider/main.go
	go build -o cmd/spider/spider cmd/spider/main.go

.PHONY: test
test:
	go test ./...

lint:
	golangci-lint run ./...
