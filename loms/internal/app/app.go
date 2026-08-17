package http

import (
	"net/http"
	"route256/loms/internal/config"
	"route256/loms/internal/handlers"
	"route256/loms/internal/repository"
	"route256/loms/internal/services"
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
	stocksProvider := repository.NewDumpRepo()
	stocksHandler := handlers.NewStocksHandler(services.NewStocksService(stocksProvider))

	http.HandleFunc("/stocks", stocksHandler.Handle)

	return http.ListenAndServe(a.config.Addr, nil)
}
