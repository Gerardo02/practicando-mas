run: build
	@./bin/shomper

build:
	@go build -o bin/shomper cmd/main.go
