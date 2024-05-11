package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func Connect(connectURL string) error {
	conn, err := pgx.Connect(context.Background(), connectURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	return nil
}
