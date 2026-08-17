package http

import (
	"log"
	"net/http"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/products"
	"route256/cart/internal/config"

	hitem "route256/cart/internal/handlers/item"

	sitem "route256/cart/internal/services/item"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (a App) Run() error {
	lomsClient, err := loms.New("loms client", a.config.LomsAddr)
	if err != nil {
		log.Fatal(err)
	}

	productClient, err := products.New("product client", a.config.ProductAddr)
	if err != nil {
		log.Fatal(err)
	}

	itemAddHandler := hitem.NewAddHandler(sitem.NewAddService(lomsClient, productClient))

	http.HandleFunc("/cart/item/add", itemAddHandler.Handle)

	return http.ListenAndServe(a.config.Addr, nil)
}
