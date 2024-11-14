export Env := "dev"

db-up:
    docker compose up db -d

db-down:
    docker compose down db

migrate:
    go run cmd/migrate/main.go

dev:
    go run cmd/web/main.go
