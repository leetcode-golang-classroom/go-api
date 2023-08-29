.PHONY: build

build:
	@go build -o ./bin/server cmd/api/main.go

run: build
	@./bin/server