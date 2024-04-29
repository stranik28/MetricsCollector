package storage

import (
	"errors"
)

var (
	ErrorMetricsNotFound      = errors.New("метрика не найдена")
	ErrorIncorrectTypeInt64   = errors.New("значение должно быть в формате int64")
	ErrorIncorrectTypeFloat64 = errors.New("значение должно быть в формате float64")
	ErrorIncorrectTypeMetrics = errors.New("недопустимый тип метрики")
)

type Metric struct {
	Gauge   float64
	Counter int64
}

type MemStorage struct {
	metrics map[string]Metric
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]Metric),
	}
}

var storage = NewMemStorage()

func SetMemStorage(key string, value Metric) {
	storage.metrics[key] = value
}

func GetMemStorage(key string) (Metric, bool) {
	value, err := storage.metrics[key]
	return value, err
}

func GetAll() (map[string]Metric, error) {
	return storage.metrics, nil
}

func ClearStorage() {
	storage = NewMemStorage()
}

func SetMemStorageMetric(metrics map[string]Metric) {
	storage.metrics = metrics
}
