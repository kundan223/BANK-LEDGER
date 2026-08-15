package main

import (
	"context"
	"fmt"
	"net/http"

	"bank-ledger/internal/api"
	"bank-ledger/internal/config"
	"bank-ledger/internal/db"
	"bank-ledger/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		fmt.Println("Failed to create pool:", err)
		return
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return
	}
	store := db.NewStore(pool)
	userService := service.NewUserService(store)
	server := api.NewServer(store, userService)

	fmt.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", server.Handler())
	if err != nil {
		fmt.Println("Server error:", err)
	}
}