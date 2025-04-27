package main

import (
	"fmt"
	"go-lessons/junior-rest-api/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // Путь к файлу .env
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	cfg := config.LoadConfig()
	fmt.Println(cfg)
}
