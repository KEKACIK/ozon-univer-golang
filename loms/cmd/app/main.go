package main

import (
	"log"

	httpapp "github.com/KEKACIK/ozon-univer-golang/loms/internal/app"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
)

func main() {
	cfg := config.NewConfigFromFlags()

	app := httpapp.NewApp(cfg)

	log.Fatal(app.Run())
}
