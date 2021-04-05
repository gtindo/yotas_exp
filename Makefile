#!make
-include .env

PG_STRING = 'host=${PG_HOST} dbname=${PG_DBNAME} user=${PG_USER} password=${PG_PASSWORD} sslmode=disable'

build:
	go build -o bin/main main.go

goose-up:
	cd migrations && goose postgres $(PG_STRING) up 

goose-down:
	cd migrations && goose postgres $(PG_STRING) down

goose-status:
	cd migrations && goose postgres $(PG_STRING) status

run:
	go run main.go