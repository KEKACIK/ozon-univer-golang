package main

import (
	"log"

	httpapp "github.com/KEKACIK/ozon-univer-golang/cart/internal/app"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/config"
)

func main() {
	cfg := config.NewConfigFromFlags()

	app := httpapp.NewApp(cfg)

	log.Fatal(app.Run())
}
