package app

import "os"

type Env struct {
	PgHost          string
	PgDbName        string
	PgUser          string
	PgPassword      string
	PgPort          string
	Port            string
	BaseUri         string
	GithubClientId  string
	GithubSecretKey string
}

var env *Env

func InitEnv() {
	env = &Env{
		PgHost:          os.Getenv("PG_HOST"),
		PgDbName:        os.Getenv("PG_DBNAME"),
		PgUser:          os.Getenv("PG_USER"),
		PgPassword:      os.Getenv("PG_PASSWORD"),
		PgPort:          os.Getenv("PG_PORT"),
		BaseUri:         os.Getenv("BASE_URI"),
		GithubClientId:  os.Getenv("GITHUB_CLIENT_ID"),
		GithubSecretKey: os.Getenv("GITHUB_SECRET_KEY"),
	}
}
