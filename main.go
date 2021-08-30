package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/betelgeuse-7/twitt/db"
	dotenv "github.com/joho/godotenv"
)

const PORT = ":8000"

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Println("Couldn't load env variables. ERR -> ", err)
	}
	postgres := db.Postgres{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
	db, err := postgres.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	r := routes()

	log.Println(fmt.Sprintf("Starting server at localhost%s\n", PORT))
	log.Fatalln(http.ListenAndServe(PORT, r))
}
