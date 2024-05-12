package service

import (
	"github.com/stranik28/MetricsCollector/internal/server/models"
	"github.com/stranik28/MetricsCollector/internal/server/storage"
)

func UpdateMetrics(reqModel models.Metrics) (models.Metrics, error) {
	switch reqModel.MType {
	case "counter":
		val, ok := storage.GetMemStorage(reqModel.ID)
		if !ok {
			val = storage.Metric{
				Gauge:   0,
				Counter: *reqModel.Delta,
			}
		} else {
			val.Counter += *reqModel.Delta
		}
		storage.SetMemStorage(reqModel.ID, val)
		reqModel.Value = &val.Gauge
	case "gauge":
		val, ok := storage.GetMemStorage(reqModel.ID)
		if !ok {
			val = storage.Metric{
				Gauge:   *reqModel.Value,
				Counter: 0,
			}
		} else {
			val.Gauge = *reqModel.Value
		}
		storage.SetMemStorage(reqModel.ID, val)
		reqModel.Value = &val.Gauge
	default:
		err := storage.ErrorIncorrectTypeMetrics
		return reqModel, err
	}
	return reqModel, nil
}
