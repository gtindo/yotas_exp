#!make
include .env
export $(shell sed 's/=.*//' .env)

PG_STRING = 'dbname=${PG_DBNAME} user=${PG_USER} password=${PG_PASSWORD}'

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