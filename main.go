package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
