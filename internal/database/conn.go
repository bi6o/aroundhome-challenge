package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driverPSQL = "postgres"
)

func Connect(ctx context.Context, dbURI string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driverPSQL, dbURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}
