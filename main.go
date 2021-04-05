package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gtindo/yotas_exp/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app.InitEnv()

	app.InitDb()

	httpPort := os.Getenv("PORT")
	defer http.ListenAndServe(":"+httpPort, app.GetRouter())

	log.Println("HTTP Server Started on port ", httpPort)
}
