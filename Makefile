include .env
export

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/cinema.go

run:
	go run ./cmd/app

test: ### run test
	go test -v ./...