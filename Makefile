dev:
	go run cmd/main.go
migrate:
	go run migration/main.go
dbstart:
	docker compose up -d
dbstop:
	docker compose down				