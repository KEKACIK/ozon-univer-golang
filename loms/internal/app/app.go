package http

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/KEKACIK/ozon-univer-golang/loms/internal/api"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/repository"
	"github.com/KEKACIK/ozon-univer-golang/loms/internal/services"
	desc "github.com/KEKACIK/ozon-univer-golang/loms/pkg/api/loms/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	// httpapp "github.com/KEKACIK/ozon-univer-golang/loms/internal/app"
	// "github.com/KEKACIK/ozon-univer-golang/loms/internal/config"
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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer( // grpc сервер (aka http.Serever)
	// grpc.ChainUnaryInterceptor( // Unary интерсепторы (aka middleware)
	// 	panic.Interceptor,
	// 	logging.Interceptor,
	// ),
	// grpc.ChainStreamInterceptor( // Stream интерсепторы (aka middleware)
	// // func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	// // },
	// ),
	)

	reflection.Register(grpcServer) // Рефлексия! (Повзоляет получать описание rpc функционала нашего сервиса. Полезно для Postman)

	provider := repository.NewDumpRepo()

	controller := api.NewHandler(services.NewStocksService(provider))

	desc.RegisterLomsServer(grpcServer, controller) // Вешаем наш обработчик (controller) на grpc сервер

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = grpcServer.Serve(lis); err != nil { // запускаем grpc сервер
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Создаем коннект с grpc сервером
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
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

	gwServer := &http.Server{ // Создаем HTTP gateway сервер
		Addr: fmt.Sprintf(":%d", a.config.HTTPPort),
		// Handler: logging.WithHTTPLoggingMiddleware(mux), // middleware
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on %d\n", a.config.GRPCPort) // запускаем HTTP сервер

	return gwServer.ListenAndServe()

	// provider := repository.NewDumpRepo()

	// // Orders
	// orderCreateHandler := ohandler.NewCreateHandler(sorders.NewCreateService(provider))
	// orderInfoHandler := ohandler.NewInfoHandler(sorders.NewInfoService(provider))
	// orderPayHandler := ohandler.NewPayHandler(sorders.NewPayService(provider))
	// orderCancelHandler := ohandler.NewCancelHandler(sorders.NewCancelService(provider))
	// // Stocks
	// // stocksHandler := handlers.NewStocksHandler(services.NewStocksService(provider))

	// http.HandleFunc("/order/create", orderCreateHandler.Handle)
	// http.HandleFunc("/order/info", orderInfoHandler.Handle)
	// http.HandleFunc("/order/pay", orderPayHandler.Handle)
	// http.HandleFunc("/order/cancel", orderCancelHandler.Handle)
	// // http.HandleFunc("/stocks", stocksHandler.Handle)

	// http.HandleFunc("/provider", provider.Test) // TODO: testing

	// fmt.Printf("App starting %s\n", a.config.Addr)
	// return http.ListenAndServe(a.config.Addr, nil)
}
