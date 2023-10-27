include .env

migrate_up:
	goose -dir ./migrations postgres "user=postgres password=$(POSTGRES_PASSWORD) dbname=postgres sslmode=disable" up

docker:
	docker compose up

run:
	go run app/main.go
