package main

import (
	"effective-task/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	if err := server.NewServer(); err != nil {
		log.Fatal(err)
	}
}
