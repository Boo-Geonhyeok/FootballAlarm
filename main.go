package main

import (
	"log"

	database "github.com/Boo-Geonhyeok/football-alarm/Database"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.InitDB()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
