package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/storage"
	"strconv"
	"strings"
)

func UpdateMetrics(metricType string, metricValue string, metricName string) error {
	switch metricType {
	case "counter":
		value, err := strconv.ParseInt(strings.TrimSpace(metricValue), 10, 64)
		if err != nil {
			err = storage.ErrorIncorrectTypeInt64
			return err
		}
		val, ok := storage.GetMemStorage(metricName)
		if !ok {
			val = storage.Metric{
				Gauge:   0,
				Counter: value,
			}
		} else {
			val.Counter += value
		}
		storage.SetMemStorage(metricName, val)
	case "gauge":
		value, err := strconv.ParseFloat(strings.TrimSpace(metricValue), 64)
		if err != nil {
			err = storage.ErrorIncorrectTypeFloat64
			return err
		}
		val, ok := storage.GetMemStorage(metricName)
		if !ok {
			val = storage.Metric{
				Gauge:   value,
				Counter: 0,
			}
		} else {
			val.Gauge = value
		}
		storage.SetMemStorage(metricName, val)
	default:
		err := storage.ErrorIncorrectTypeMetrics
		return err
	}
	return nil
}
