package storage

import (
	"context"
	"github.com/stranik28/MetricsCollector/internal/server"
	"os"
	"time"
)

func InitSaveMem(filename string, restore bool, interval int, done chan os.Signal) {
	c := context.Background()
	db, _ := NewDBConnection(c, server.DBURL)
	go func() {
		<-done
		if db == nil {
			SaveMetricsToFile(server.FileStoragePath)
		}
		os.Exit(0)
	}()

	periodDuration := time.Duration(interval) * time.Second
	if restore {
		var metrics map[string]Metric
		var err error
		if db == nil {
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
		if server.DBURL == "" {
			SaveMetricsToFile(filename)
		}
	}
}
