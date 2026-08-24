package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/api"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/services"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/services/orders"
	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config *config.Config
	DBPool *pgxpool.Pool
}

func NewApp(config *config.Config, dbPool *pgxpool.Pool) *App {

	return &App{
		config: config,
		DBPool: dbPool,
	}
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	controller := api.NewHandler(
		orders.NewCreateService(a.DBPool),
		orders.NewInfoService(a.DBPool),
		orders.NewPayService(a.DBPool),
		orders.NewCancelService(a.DBPool),

		services.NewStocksService(a.DBPool),
	)

	desc.RegisterLomsServer(grpcServer, controller)

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		lis.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()

	err = desc.RegisterLomsHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.HTTPPort),
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on %d\n", a.config.HTTPPort) // запускаем HTTP сервер

	return gwServer.ListenAndServe()
}
