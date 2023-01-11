package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TEST")

}
