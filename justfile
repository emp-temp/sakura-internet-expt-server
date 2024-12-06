export Env := "dev"

login-server:
    sshpass -pZvXe6U7sKbms ssh ubuntu@163.43.144.159

db-up:
    docker compose up db -d

db-down:
    docker compose down db

migrate:
    go run cmd/migrate/main.go

dev:
    go run cmd/web/main.go
