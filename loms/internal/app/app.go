package http

import (
	"net/http"
	"route256/loms/internal/config"
	"route256/loms/internal/handlers"
	ohandler "route256/loms/internal/handlers/orders"
	"route256/loms/internal/repository"
	"route256/loms/internal/services"
	sorders "route256/loms/internal/services/orders"
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
	http.HandleFunc("/stocks", orderCancelHandler.Handle)
	http.HandleFunc("/stocks", stocksHandler.Handle)

	http.HandleFunc("/provider", provider.Test) // TODO: testing

	return http.ListenAndServe(a.config.Addr, nil)
}
