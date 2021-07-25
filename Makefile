.PHONY: build

start: up
	go build -v ./cmd/api
	./api

up:
	docker-compose up -d --build

migrate-up:
	migrate --path migrations -database "postgres://localhost:54300/godb?sslmode=disable&user=godb&password=godb" up
migrate-down:
	migrate --path migrations -database "postgres://localhost:54300/godb?sslmode=disable&user=godb&password=godb" down
migrate-create:
	migrate create -ext sql -dir migrations UsersCreationsMigrations

.DEFAULT_GOAL:= build