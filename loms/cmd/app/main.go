package main

import (
	"log"
	httpapp "route256/loms/internal/app"
	"route256/loms/internal/config"
)

func main() {
	cfg := config.NewConfigFromFlags()

	app := httpapp.NewApp(cfg)

	log.Fatal(app.Run())
}
