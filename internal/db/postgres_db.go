package db

import (
	"context"
	"equi_genea_account_service/config"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewPostgresDB(cfg *config.DatabaseConfig) (*Queries, *pgx.Conn, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, nil, err
	}

	queries := New(conn)

	return queries, conn, nil
}
