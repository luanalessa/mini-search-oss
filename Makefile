SHELL := /bin/bash

.PHONY: up down logs build run lint test tidy

up:
	docker-compose up -d --build

down:
	docker compose down -v

logs:
	docker compose logs -f --tail=100

build:
	docker build -t minisearch:dev .

run:
	go run ./cmd/api

lint:
	golangci-lint run

test:
	go test ./... -v

tidy:
	go mod tidy
