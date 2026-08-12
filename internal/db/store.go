package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store struct {
	*Queries
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: New(pool),
	}
}
