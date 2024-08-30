include .env
export

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database $(DB_URL) up

.PHONY: migrate_down	
migrate_down:
	migrate -path migrations -database $(DB_URL) down

.DEFAULT_GOAL := build