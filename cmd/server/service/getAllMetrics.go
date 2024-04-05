package service

import "github.com/stranik28/MetricsCollector/cmd/server/storage"

func GetAllMetrics() (map[string]storage.Metric, error) {
	metrics, err := storage.GetAll()
	return metrics, err
}