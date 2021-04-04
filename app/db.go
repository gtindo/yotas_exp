package app

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	var err error

	params := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.PgHost, env.PgUser, env.PgPassword, env.PgDbName, env.PgPort,
	)

	db, err = gorm.Open(postgres.Open(params), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL server", err)
	}

	log.Println("Conntected to PostgreSQL Database")
}
