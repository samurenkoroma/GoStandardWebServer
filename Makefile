.PHONY: build

start: up
	go build -v ./cmd/api
	./api

up:
	docker-compose up -d --build

.DEFAULT_GOAL:= build