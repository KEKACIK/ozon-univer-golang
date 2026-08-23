package main

import (
	"log"

	myApp "github.com/KEKACIK/ozon-univer-golang/loms/internal/app"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
)

func main() {
	cfg := config.NewConfigFromYaml()

	app := myApp.NewApp(cfg)

	log.Fatal(app.Run())
}
