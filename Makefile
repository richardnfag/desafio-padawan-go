.PHONY: test build run dotenv coverage make_migrations create_migration migrate
test:
	go test -v ./...

build:
	go build -o bin/ ./...

run:
	go run cmd/main.go

dotenv:
	cp .env.sample .env

coverage:
	go test -v -coverpkg=./... -coverprofile=coverage/coverage.out ./...
	go tool cover -func coverage/coverage.out

make_migrations:
	atlas migrate diff --env gorm
	atlas migrate hash

create_migration:
	atlas migrate new $(name)
	atlas migrate hash

migrate:
	atlas migrate apply --url "mysql://exchange:exchange@db:3306/exchange_db" --dir "file://migrations"
