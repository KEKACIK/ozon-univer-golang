package main

import (
	"log"

	myApp "github.com/KEKACIK/ozon-univer-golang/cart/internal/app"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/config"
)

func main() {
	cfg := config.NewConfigFromFlags()

	app := myApp.NewApp(cfg)

	log.Fatal(app.Run())
}
