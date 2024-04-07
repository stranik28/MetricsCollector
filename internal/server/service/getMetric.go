package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func GetMetricByName(metricName string, metricType string) (interface{}, error) {
	metric, ok := storage.GetMemStorage(metricName)
	if !ok {
		err := storage.ErrorMetricsNotFound
		return nil, err
	}
	if metricType == "gauge" {
		return metric.Gauge, nil
	} else if metricType == "counter" {
		return metric.Counter, nil
	} else {
		err := storage.ErrorMetricsNotFound
		return nil, err
	}

}
