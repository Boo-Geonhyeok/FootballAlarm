package main

import (
	"log"

	api "github.com/Boo-Geonhyeok/football-alarm/API"
	database "github.com/Boo-Geonhyeok/football-alarm/Database"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.InitDB()
	// api.GetAllFixtures()
	api.GetStanding(39)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
