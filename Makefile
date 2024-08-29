include .env

run: build
	@./bin/shomper

down: 
	@cd cmd/schema && goose postgres ${DB_URL} down

up: 
	@cd cmd/schema && goose postgres ${DB_URL} up

build:
	@go build -o bin/shomper cmd/main.go

sqlgen:
	@sqlc generate
