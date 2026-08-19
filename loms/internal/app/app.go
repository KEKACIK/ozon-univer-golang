package http

import (
	"fmt"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/handlers"
	ohandler "github.com/KEKACIK/ozon-univer-golang/loms/internal/handlers/orders"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/services"
	sorders "github.com/KEKACIK/ozon-univer-golang/loms/internal/services/orders"
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
	provider := repository.NewDumpRepo()

	// Orders
	orderCreateHandler := ohandler.NewCreateHandler(sorders.NewCreateService(provider))
	orderInfoHandler := ohandler.NewInfoHandler(sorders.NewInfoService(provider))
	orderPayHandler := ohandler.NewPayHandler(sorders.NewPayService(provider))
	orderCancelHandler := ohandler.NewCancelHandler(sorders.NewCancelService(provider))
	// Stocks
	stocksHandler := handlers.NewStocksHandler(services.NewStocksService(provider))

	http.HandleFunc("/order/create", orderCreateHandler.Handle)
	http.HandleFunc("/order/info", orderInfoHandler.Handle)
	http.HandleFunc("/order/pay", orderPayHandler.Handle)
	http.HandleFunc("/order/cancel", orderCancelHandler.Handle)
	http.HandleFunc("/stocks", stocksHandler.Handle)

	http.HandleFunc("/provider", provider.Test) // TODO: testing

	fmt.Printf("App starting %s\n", a.config.Addr)
	return http.ListenAndServe(a.config.Addr, nil)
}
