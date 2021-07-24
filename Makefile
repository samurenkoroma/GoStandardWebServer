.PHONY: build

start:
	go build -v ./cmd/api
	./api

.DEFAULT_GOAL:= build