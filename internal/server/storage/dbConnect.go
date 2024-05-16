package storage

import (
	"context"
	"database/sql"
	"github.com/jackc/pgerrcode"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"time"
)

type DBConnection struct {
	Conn *sql.DB
	url  string
}

func NewDBConnection(ctx context.Context, url string) (*DBConnection, error) {
	dbCon, err := connect(ctx, url)
	if err != nil {
		return nil, err
	}
	return &DBConnection{
		url:  url,
		Conn: dbCon,
	}, nil
}

func connect(ctx context.Context, connectURL string) (*sql.DB, error) {
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
