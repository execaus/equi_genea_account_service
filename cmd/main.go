package main

import (
	"context"
	"equi_genea_account_service/config"
	"equi_genea_account_service/internal/db"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbConn, err := db.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbConn.Close(context.Background())

	fmt.Println("Connected to DB successfully!")
}
