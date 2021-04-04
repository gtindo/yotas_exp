package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gtindo/yotas_exp/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.InitEnv()

	app.InitDb()

	httpPort := os.Getenv("PORT")
	defer http.ListenAndServe(":"+httpPort, app.GetRouter())

	log.Println("HTTP Server Started on port ", httpPort)
}
