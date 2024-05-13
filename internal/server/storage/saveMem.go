package storage

import (
	"context"
	"database/sql"
	"github.com/stranik28/MetricsCollector/internal/server"
	"os"
	"time"
)

func InitSaveMem(filename string, restore bool, interval int, done chan os.Signal) {
	c := context.Background()
	db, _ := Connect(c, server.DBURL)
	go func(db *sql.DB) {
		<-done
		if db != nil {
			err := saveMetricsToDB(c, db)
			if err != nil {
				SaveMetricsToFile(server.FileStoragePath)
			}
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
			metrics, err = loadMetricsFromDB(c, db)
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
			err := saveMetricsToDB(c, db)
			if err != nil {
				SaveMetricsToFile(server.FileStoragePath)
			}
		} else {
			SaveMetricsToFile(filename)
		}
	}
}
