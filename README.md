# Yotas Explorations


## Purpose

This project Was creating for testing Github Oauth API, GORM and goose.

[Github OAUTH API](https://docs.github.com/en/developers/apps/authorizing-oauth-apps#troubleshooting)
[GORM](https://gorm.io/)
[Goose](https://github.com/pressly/goose)

![illustration](/yotas_exp.gif)

## How to run

First create the .env file using env.example as template

### With sources

#### Prerequisite
- Golang 1.14+
- Make
- PostgreSQL 12+

#### Get dependencies
```shell
go get -d -v ./...
```

#### Build
```shell
make build
```

#### Run migrations
```shell
make goose-up
```

#### Run
```shell
make run
```

### With docker

```shell
docker-compose build
docker-compose up
```

## Author
[Gtindo Dev](https://gtindo.dev)