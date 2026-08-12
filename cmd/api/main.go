package main

import (
	"context"
	"fmt"

	"bank-ledger/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(
		ctx,
		"postgresql://postgres:postgres@localhost:5432/bank_ledger?sslmode=disable",
	)

	if err != nil {
		fmt.Println("Failed to create pool:", err)
		return
	}

	defer pool.Close()

	err = pool.Ping(ctx)

	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL:", err)
		return
	}

	store := db.NewStore(pool)

	fmt.Println("Connected to PostgreSQL")
	fmt.Println(store)
}