package storage

import (
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/server"
	"os"
	"time"
)

func SaveMetricsToFile(filename string) {
	metrics, err := GetAll()
	if err != nil {
	}
	data, err := json.Marshal(metrics)
	if err != nil {
		return
	}
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
	}
}

func LoadMetricsFromFile(filename string) (map[string]Metric, error) {
	var metrics map[string]Metric
	data, err := os.ReadFile(filename)
	if err != nil {
		return metrics, err
	}
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		if len(data) == 0 {
			return make(map[string]Metric), nil
		}
		return metrics, err
	}
	return metrics, nil
}

func InitFileSave(filename string, restore bool, interval int, done chan os.Signal) {
	go func() {
		<-done
		SaveMetricsToFile(server.FileStoragePath)
		os.Exit(0)
	}()

	periodDuration := time.Duration(interval) * time.Second
	if restore {
		metrics, err := LoadMetricsFromFile(server.FileStoragePath)
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
		SaveMetricsToFile(filename)
	}

}
