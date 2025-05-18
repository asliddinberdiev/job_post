.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

run: build
	docker-compose up --remove-orphans app

swag:
	swag init -g internal/handler/handler.go -o docs --parseVendor --parseDependency --parseInternal --parseDepth 1
