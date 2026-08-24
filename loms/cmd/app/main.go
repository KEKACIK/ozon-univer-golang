package main

import (
	"context"
	"fmt"
	"log"
	"os"

	myApp "github.com/KEKACIK/ozon-univer-golang/loms/internal/app"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.NewConfigFromYaml()

	dbPool, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	app := myApp.NewApp(cfg, dbPool)

	log.Fatal(app.Run())
}
