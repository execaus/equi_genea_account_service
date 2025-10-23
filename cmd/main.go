package main

import (
	"context"
	"equi_genea_account_service/config"
	"equi_genea_account_service/internal/app"
	"equi_genea_account_service/internal/db"
	accountpb "equi_genea_account_service/internal/pb/api/account"
	"equi_genea_account_service/internal/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	queries, dbConn, err := db.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	fmt.Println("Connected to DB successfully!")

	grpcServer := grpc.NewServer()
	accountService := service.NewAccountService(queries)
	handler := app.NewAccountHandler(accountService)

	accountpb.RegisterAccountServiceServer(grpcServer, handler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go runServer(grpcServer, cfg)

	<-stop
	fmt.Println("\nShutting down gracefully...")
	grpcServer.GracefulStop()
	err = dbConn.Close(context.Background())
	if err != nil {
		log.Fatalf("Faled to close connect to db: %s", err.Error())
	}
	fmt.Println("Server stopped")
}

func runServer(grpcServer *grpc.Server, cfg *config.Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
	fmt.Printf("Starting gRPC server on port %s...\n", cfg.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
