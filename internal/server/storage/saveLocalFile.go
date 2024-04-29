package storage

import (
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/server"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"go.uber.org/zap"
	"os"
	"time"
)

func SaveMetricsToFile(filename string) {
	metrics, err := GetAll()
	if err != nil {
		logger.Log.Error("Error getting metrics from file", zap.Error(err))
	}
	logger.Log.Debug("Saving metrics to file", zap.String("filename", filename))
	data, err := json.Marshal(metrics)
	if err != nil {
		logger.Log.Error("Ошибка маршалинга метрик:", zap.Any("Error", err))
		return
	}
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		logger.Log.Error("Ошибка сохранения метрик в файл:", zap.Any("Error", err))
	}
}

func LoadMetricsFromFile(filename string) (map[string]Metric, error) {
	var metrics map[string]Metric
	logger.Log.Debug("Loading metrics from file", zap.String("filename", filename))
	data, err := os.ReadFile(filename)
	if err != nil {
		return metrics, err
	}
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		return metrics, err
	}
	return metrics, nil
}

func InitFileSave(filename string, restore bool, interval int, done chan os.Signal) {

	go func() {
		<-done
		logger.Log.Info("Received signal to exit. Saving metrics and shutting down.")
		SaveMetricsToFile(server.FileStoragePath)
		os.Exit(0)
	}()

	periodDuration := time.Duration(interval) * time.Second
	logger.Log.Debug("InitFileSave")
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
