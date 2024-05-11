package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(connectURL string, ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectURL)
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
