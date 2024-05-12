package storage

import (
	"context"
	"database/sql"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"time"
)

func Connect(connectURL string, ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectURL)
	if err != nil {
		return nil, err
	}
	err = db.PingContext(ctx)
	if err != nil {
		retries := []int{1, 3, 5}
		for _, v := range retries {
			if err := db.PingContext(ctx); err != nil {
				if pgerrcode.IsConnectionException(err.Error()) {
					time.Sleep(time.Duration(v) * time.Second)
				} else {
					return nil, err
				}
			} else {
				return db, nil
			}
		}
		return nil, err
	}
	return db, nil
}
