package main

import (
	"log"
	"net/http"
	"route256/loms/internal/handlers"
	"route256/loms/internal/repository"
	"route256/loms/internal/services"
)

func main() {
	stocksProvider := repository.NewDumpRepo()
	stocksHandler := handlers.NewStocksHandler(services.NewStocksService(stocksProvider))

	http.HandleFunc("/stocks", stocksHandler.Handle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
