package main

import (
	application "api-service"
	"api-service/internal/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()
	application.StartService(cfg)
}
