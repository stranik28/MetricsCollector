package storage

import (
	"encoding/json"
	"github.com/stranik28/MetricsCollector/internal/server/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// saveMetricsToFile сохраняет метрики в файл
func saveMetricsToFile(filename string) {
	metrics, err := GetAll()
	if err != nil {
		logger.Log.Error("Error getting metrics from file", zap.Error(err))
	}
	logger.Log.Info("Saving metrics to file", zap.String("filename", filename))
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

// loadMetricsFromFile загружает метрики из файла
func loadMetricsFromFile(filename string) (map[string]Metric, error) {
	var metrics map[string]Metric
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

func saveMetricsPeriodically(filename string, period time.Duration) {
	logger.Log.Info("Init Save Metrics Periodically", zap.String("filename", filename))
	for {
		logger.Log.Info("Saving metrics")

		// Сохраняем метрики
		saveMetricsToFile(filename)

		// Ждем определенное время перед следующим сохранением
		time.Sleep(period * time.Second)
	}
}

func InitFileSave(filename string, load bool, interval time.Duration) {
	logger.Log.Info("InitFileSave")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	logger.Log.Info("InitDoneSignal")

	logger.Log.Info("InitMetricsChan")

	if load {
		logger.Log.Info("Load Metrics")               // объявление err в этом контексте
		metrics, err := loadMetricsFromFile(filename) // используем оригинальную переменную err
		if err != nil {
			logger.Log.Warn("Ошибка загрузки метрик:", zap.Any("Error", err))
		}
		storage = &MemStorage{metrics: metrics}
	}

	// Горутина для сохранения метрик с заданной периодичностью
	go saveMetricsPeriodically(filename, interval)

	// Горутина для обработки сигналов завершения
	go func() {
		<-done
		// При получении сигнала завершения сохраняем текущие метрики
		saveMetricsToFile(filename)
		os.Exit(1)
	}()

}
