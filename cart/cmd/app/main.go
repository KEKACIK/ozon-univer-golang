package main

import (
	"log"
	"net/http"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/products"
	hitem "route256/cart/internal/handlers/item"
	sitem "route256/cart/internal/services/item"
)

func main() {
	lomsClient, err := loms.New("loms client", "loms:8080")
	if err != nil {
		log.Fatal(err)
	}

	// ProductClient к сожалению не работает на 08.2026
	// Поэтому всё захардкожено (харкодено хз как правильно)
	productClient, err := products.New("product client", "HARDCORREEE")
	if err != nil {
		log.Fatal(err)
	}

	itemAddHandler := hitem.NewAddHandler(sitem.NewAddService(lomsClient, productClient))

	http.HandleFunc("cart/item/add", itemAddHandler.Handle)
}
