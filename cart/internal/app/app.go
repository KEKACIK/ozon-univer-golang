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
	"github.com/KEKACIK/ozon-univer-golang/cart/internal/services/item"
	desc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/cart/v1"
	ldesc "github.com/KEKACIK/ozon-univer-golang/cart/pkg/api/loms/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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
		item.NewAddService(lomsClient, productClient),
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
