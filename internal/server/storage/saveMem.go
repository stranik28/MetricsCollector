package storage

import (
	"context"
	"database/sql"
	"github.com/stranik28/MetricsCollector/internal/server"
	"os"
	"time"
)

func InitSaveMem(filename string, restore bool, interval int, done chan os.Signal) {
	db, _ := Connect(server.DBURL, context.Background())
	go func(db *sql.DB) {
		<-done
		if db != nil {
			saveMetricsToDB(db)
		} else {
			SaveMetricsToFile(server.FileStoragePath)
		}
		os.Exit(0)
	}(db)

	periodDuration := time.Duration(interval) * time.Second
	if restore {
		var metrics map[string]Metric
		var err error
		if db != nil {
			metrics, err = loadMetricsFromDB(db)
		} else {
			metrics, err = LoadMetricsFromFile(server.FileStoragePath)
		}
		if err != nil {
			return
		}
		if metrics != nil {
			SetMemStorageMetric(metrics)
		}
	}

	// Запускаем таймер для сохранения метрик на диск с указанной периодичностью
	ticker := time.NewTicker(periodDuration)
	defer ticker.Stop()

	for range ticker.C {
		if server.DBURL != "" {
			saveMetricsToDB(db)
		} else {
			SaveMetricsToFile(filename)
		}
	}
}
