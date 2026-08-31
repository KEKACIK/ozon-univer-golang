package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/cart/internal/api"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/clients/loms"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/clients/products"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/config"
	service "github.com/KEKACIK/ozon-univer-golang/cart/internal/services"
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/services/item"
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
	ldesc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/loms/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config *config.Config
	dbPool *pgxpool.Pool
}

func NewApp(config *config.Config, dbPool *pgxpool.Pool) *App {

	return &App{
		config: config,
		dbPool: dbPool,
	}
}

func (a *App) Run() error {
	lomsConn, err := grpc.NewClient(
		a.config.LomsGRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer lomsConn.Close()

	grpcLomsClient := ldesc.NewLomsServiceClient(lomsConn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	lomsClient := loms.NewClient(grpcLomsClient) // TODO: a.config.LomsAddr)
	productClient := products.NewClient()        // TODO: a.config.ProductAddr)

	controller := api.NewHandler(
		service.NewListService(a.dbPool, productClient),
		service.NewClearService(a.dbPool),
		service.NewCheckoutService(a.dbPool, lomsClient),
		item.NewAddService(a.dbPool, lomsClient, productClient),
		item.NewDeleteService(a.dbPool, productClient),
	)

	desc.RegisterCartServiceServer(grpcServer, controller)

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

	err = desc.RegisterCartServiceHandler(context.Background(), mux, conn)
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
