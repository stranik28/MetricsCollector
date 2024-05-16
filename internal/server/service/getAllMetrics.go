package service

import (
	"context"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func GetAllMetrics(c context.Context) (map[string]storage.Metric, error) {
	if server.DBURL != "" {
		db, err := storage.NewDBConnection(c, server.DBURL)
		if err != nil {
			return nil, err
		}
		metrics, err := storage.LoadMetricsFromDB(c, db)
		return metrics, err
	}
	metrics, err := storage.GetAll()
	return metrics, err
}
