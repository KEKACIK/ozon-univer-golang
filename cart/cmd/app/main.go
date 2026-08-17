package main

import (
	"log"
	httpapp "route256/cart/internal/app"
	"route256/cart/internal/config"
)

func main() {
	cfg := config.NewConfigFromFlags()

	app := httpapp.NewApp(cfg)

	log.Fatal(app.Run())
}
