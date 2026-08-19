package http

import (
	"log"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/clients/loms"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/clients/products"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/config"

	hitem "github.com/KEKACIK/ozon-univer-golang/cart/internal/handlers/item"

	sitem "github.com/KEKACIK/ozon-univer-golang/cart/internal/services/item"
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
