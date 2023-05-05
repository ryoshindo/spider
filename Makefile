.PHONY: build
build: cmd/spider/main.go
	go build -o cmd/spider/spider cmd/spider/main.go

.PHONY: test
test:
	go test ./...

lint:
	golangci-lint run ./...

test-apply: cmd/spider/spider
	make build
	./cmd/spider/spider apply --config ../spider-test/spider.yml

test-destroy: cmd/spider/spider
	make build
	./cmd/spider/spider destroy --config ../spider-test/spider.yml
