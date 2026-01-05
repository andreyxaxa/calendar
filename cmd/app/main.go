package main

import (
	"log"

	"github.com/andreyxaxa/calendar/config"
	"github.com/andreyxaxa/calendar/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file found")
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	app.Run(cfg)
}
